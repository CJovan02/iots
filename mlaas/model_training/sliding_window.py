import numpy as np

def create_sliding_window(data: tuple, window_size: int, future_prediction: int) -> tuple[np.ndarray, np.ndarray]:
    x = np.array(data[0])
    y = np.array(data[1]).reshape(-1)

    x_win, y_win = [], []
    for i in range(len(x) - window_size - future_prediction):
        x_win.append(x[i: i + window_size])
        y_win.append(y[i + window_size + future_prediction].max())

    return np.array(x_win), np.array(y_win)