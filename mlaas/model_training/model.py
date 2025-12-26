import pandas as pd
from pandas import DataFrame, Series
from sklearn.metrics import classification_report, confusion_matrix
from sklearn.model_selection import train_test_split
from sklearn.ensemble import RandomForestClassifier
import numpy as np

def load_dataset():
    df = pd.read_csv("../data/smoke_detection_iot.csv")

    features = df.drop(
        columns=["Fire Alarm", "UTC", "Unnamed: 0", "CNT"]
    )
    labels = df["Fire Alarm"]

    x = features.values
    y = labels.values
    window_size = 50
    future_window = 200

    x_new, y_new = [], []
    for i in range(len(x) - window_size - future_window):
        x_window = x[i: i + window_size].flatten()
        y_window = y[i + window_size: i + window_size + future_window].max()
        x_new.append(x_window)
        y_new.append(y_window)

    return np.array(x_new), np.array(y_new)


x, y = load_dataset()

x_train, x_test, y_train, y_test = train_test_split(
    x, y,
    test_size=0.2,
    shuffle=False
)

model = RandomForestClassifier(n_estimators=100, random_state=42)
model.fit(x_train, y_train)

y_pred = model.predict(x_test)

print(classification_report(y_test, y_pred))
print(confusion_matrix(y_test, y_pred))
