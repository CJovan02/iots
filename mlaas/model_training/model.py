import pandas as pd

from sklearn.preprocessing import StandardScaler

from tensorflow.keras.models import Sequential
from tensorflow.keras.layers import LSTM, Dense, Input
from tensorflow.keras.optimizers import Adam
from tensorflow.keras.callbacks import ModelCheckpoint

from model_training.prediction_tests import test_predictions
from model_training.sliding_window import create_sliding_window, create_sliding_window_for_early_warning
from model_training.train_test_split import create_train_test_split

# import sys
# import numpy as np
# np.set_printoptions(threshold=sys.maxsize)

WINDOW_SIZE = 40
EARLY_WARNING_WINDOW = 500

# We don't use all the features since the server doesn't store all of them.
# This is for project simplicity reasons
features = ["Temperature[C]", "Humidity[%]", "TVOC[ppb]", "eCO2[ppm]", "Raw H2", "Raw Ethanol", "PM2.5", ]
label = "Fire Alarm"

# Load the dataset
df = pd.read_csv("smoke_detection_iot.csv")

x = df[features]
y = df[label].values

print(x.head())

# Create train_test split
train_set_raw, early_warning_test_2_raw, fire_test_raw, calm_test_raw = \
    create_train_test_split(x, y, 1000)

y = train_set_raw[0]
split_idx = int(0.8 * len(y))

val_y = y[split_idx:]
train_y_only = y[:split_idx]

# LSTM requires scaled labels
# Fit scaler only on train
scaler = StandardScaler()
train_set_raw = (
    scaler.fit_transform(train_set_raw[0]),
    train_set_raw[1]
)

# Transform test sets
early_warning_test_2_raw = (
    scaler.transform(early_warning_test_2_raw[0]),
    early_warning_test_2_raw[1]
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

# Train set includes a window before fire starts, so we pass the *fire_start_idx* in order to have a linear increase of
# y label before the fire 1
train_x, train_y = create_sliding_window_for_early_warning(train_set_raw, WINDOW_SIZE, 3177, EARLY_WARNING_WINDOW)
#print_last_20_percent((train_x, train_y))

# This test set includes only the period before fire 2 starts, so the fire start index is the last inside y array
early_warning_test = \
    create_sliding_window_for_early_warning(
        early_warning_test_2_raw,
        WINDOW_SIZE,
        len(early_warning_test_2_raw[1]) - 1,
        EARLY_WARNING_WINDOW)

# The rest of test sets don't include transitional period
fire_test = create_sliding_window(fire_test_raw, WINDOW_SIZE)
calm_test = create_sliding_window(calm_test_raw, WINDOW_SIZE)

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

# After all epochs are completed, it will look for the epoch that has the smallest "val_loss" and it will save it.
cp = ModelCheckpoint(
    filepath="../model_artifacts/model.keras",
    save_best_only=True,
    monitor="val_loss",
    mode="min",
    verbose=1
)
model.compile(optimizer=Adam(learning_rate=0.0001), loss='mse', metrics=['mae'])

# This trains the model. Validation split=0.2 grabs the last 20% of the train set to use as a validation
model.fit(
    train_x, train_y,
    epochs=50,
    batch_size=64,
    validation_split=0.2,
    shuffle=False,
    callbacks=[cp]
)

from tensorflow.keras.models import load_model

best_model = load_model("../model_artifacts/model.keras")

# Testing the model
test_predictions(best_model, early_warning_test, fire_test, calm_test)
