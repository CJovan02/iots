import pandas as pd

from matplotlib import pyplot as plt

from sklearn.metrics import classification_report
from sklearn.preprocessing import StandardScaler

from tensorflow.keras.models import Sequential
from tensorflow.keras.callbacks import ModelCheckpoint
from tensorflow.keras.layers import LSTM, Dense, Input
from tensorflow.keras.optimizers import Adam

from model_training.sliding_window import create_sliding_window
from model_training.train_test_split import create_train_test_split

WINDOW_SIZE = 70
FUTURE_PREDICTION = 200

# We don't use all the features since the server doesn't store all the features
# This is for project simplicity reasons
features = ["Temperature[C]", "Humidity[%]", "TVOC[ppb]", "eCO2[ppm]", "Raw H2", "Raw Ethanol", "PM2.5", ]
label = "Fire Alarm"

# Load the dataset
df = pd.read_csv("../data/smoke_detection_iot.csv")

x = df[features]
y = df[label].values

# y.plot()
# plt.show()

print(x.head())

# Create train_test split
train_set_raw, early_warning_test_raw, fire_test_raw, calm_test_raw = \
    create_train_test_split(x, y, 500, FUTURE_PREDICTION)

# LSTM requires scaled labels
# Fit scaler only on train
scaler = StandardScaler()
train_set_raw = (
    scaler.fit_transform(train_set_raw[0]),
    train_set_raw[1]
)

# Transform test sets
early_warning_test_raw = (
    scaler.transform(early_warning_test_raw[0]),
    early_warning_test_raw[1]
)

fire_test_raw = (
    scaler.transform(fire_test_raw[0]),
    fire_test_raw[1]
)

calm_test_raw = (
    scaler.transform(calm_test_raw[0]),
    calm_test_raw[1]
)

# Only after splitting we create sliding windows
train_x, train_y = create_sliding_window(train_set_raw, WINDOW_SIZE, FUTURE_PREDICTION)
early_war_x, early_war_y = create_sliding_window(early_warning_test_raw, WINDOW_SIZE, FUTURE_PREDICTION)
fire_test_x, fire_test_y = create_sliding_window(fire_test_raw, WINDOW_SIZE, FUTURE_PREDICTION)
calm_test_x, calm_test_y = create_sliding_window(calm_test_raw, WINDOW_SIZE, FUTURE_PREDICTION)

# Neural network
model = Sequential()
model.add(Input(shape=(WINDOW_SIZE, train_x.shape[2])))
model.add(LSTM(
    32,
    dropout=0.2,
    recurrent_dropout=0.2,
))
model.add(Dense(1, activation='sigmoid'))

print(model.summary())

cp = ModelCheckpoint("model.keras", save_best_only=True)
model.compile(optimizer=Adam(learning_rate=0.0001), loss='binary_focal_crossentropy', metrics=['accuracy'])

# Training
model.fit(train_x, train_y, epochs=10, batch_size=64, validation_split=0.2, callbacks=[cp])

# Predictions
early_war_pred = model.predict(early_war_x)
early_war_pred = (early_war_pred > 0.5).astype(int)

print(classification_report(early_war_y, early_war_pred))

early_war_results = pd.DataFrame(data={"Train Predictions": early_war_pred.flatten(), "Actual Values": early_war_y})
plt.plot(early_war_results["Train Predictions"])
plt.plot(early_war_results["Actual Values"])
plt.show()

fire_test_pred = model.predict(fire_test_x)
fire_test_pred = (fire_test_pred > 0.5).astype(int)

print(classification_report(fire_test_y, fire_test_pred))
fire_test_results = pd.DataFrame(data={"Train Predictions": fire_test_pred.flatten(), "Actual Values": fire_test_y})
plt.plot(fire_test_results["Train Predictions"])
plt.plot(fire_test_results["Actual Values"])
plt.show()

calm_test_pred = model.predict(calm_test_x)
calm_test_pred = (calm_test_pred > 0.5).astype(int)

print(classification_report(calm_test_y, calm_test_pred, zero_division=0))
calm_test_results = pd.DataFrame(data={"Train Predictions": calm_test_pred.flatten(), "Actual Values": calm_test_y})
plt.plot(calm_test_results["Train Predictions"])
plt.plot(calm_test_results["Actual Values"])
plt.show()
