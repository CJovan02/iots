import numpy as np

from app.dto import PredictionRequest


def request_to_x(request: PredictionRequest) -> np.ndarray:
    # The format for model prediction is this:
    # (batch_size, window_size, 7)

    # 7 -> number of features of the input

    # window_size -> size of the sliding window used for training the model
    # here, window_size is just the number of sensor readings being sent through request.

    # batch_size -> if we want to predict in batches of previous format of (window_size, y)
    # we don't want that, so we just add one dimension just to follow the keras model.predict format

    x = np.array([
        [
            r.temperature,
            r.humidity,
            r.tvoc,
            r.e_co2,
            r.raw_hw,
            r.raw_ethanol,
            r.pm_25
        ]
        for r in request.readings
    ])

    x = np.expand_dims(x, axis=0)
    return x

def y_to_float(y: np.ndarray) -> float:
    return y.ravel()[0].astype(float)
