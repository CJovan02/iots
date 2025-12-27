import numpy as np
import pandas as pd
from matplotlib import pyplot as plt
from sklearn.metrics import classification_report
from sklearn.model_selection import train_test_split
from sklearn.preprocessing import StandardScaler
from tensorflow.keras.layers import LSTM, Dense
from tensorflow.keras.models import Sequential
from tensorflow.python.ops.confusion_matrix import confusion_matrix

# We don't use all the features since the server doesn't store all the features
# This is for project simplicity reasons
features = ["Temperature[C]", "Humidity[%]", "TVOC[ppb]", "eCO2[ppm]", "Raw H2", "Raw Ethanol", "PM2.5", ]
label = "Fire Alarm"
1
df = pd.read_csv("../data/smoke_detection_iot.csv")

x = df[features]
y = df[label]

# y.plot()
# plt.show()

print(x.head())

scaler = StandardScaler()
x = scaler.fit_transform(x)

x_df = pd.DataFrame(x)
print(x_df.head())
print(x_df.describe())


def create_sliding_window(x, y, window_size) -> tuple[np.ndarray, np.ndarray]:
    x_win, y_win = [], []
    for i in range(len(x) - window_size):
        x_win.append(x[i: i + window_size])
        y_win.append(y[i + window_size])

    return np.array(x_win), np.array(y_win)


window_size = 30
x_win, y_win = create_sliding_window(x, y, window_size)
print(x_win.shape, y_win.shape)

x_train, x_test, y_train, y_test = train_test_split(x_win, y_win, test_size=0.2, shuffle=False)
# x_train, y_train = x[:51142], y[:51142]
# x_test, y_test = x[51142:], y[51142:]
print(x_train.shape, y_train.shape, x_test.shape, y_test.shape)

# Neural network
model = Sequential()
model.add(LSTM(
    32,
    input_shape=(window_size, x_win.shape[2])
))
model.add(Dense(1, activation='sigmoid'))

print(model.summary())

model.compile(optimizer='adam', loss='binary_crossentropy', metrics=['accuracy'])

# Training
model.fit(x_train, y_train, epochs=10, batch_size=64, validation_split=0.2)

# Evaluation
loss, acc = model.evaluate(x_test, y_test)
print(f"Test Accuracy: {acc:.4f}")

# Prediction
y_pred = model.predict(x_test)
y_pred = (y_pred > 0.5).astype(int)

print(classification_report(y_test, y_pred))
print(confusion_matrix(y_test, y_pred))
