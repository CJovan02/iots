import numpy as np

def create_sliding_window(data: tuple, window_size: int) -> tuple[
    np.ndarray, np.ndarray]:
    x = data[0]
    y = data[1]

    x_win, y_win = [], []
    for i in range(len(x) - window_size):
        x_win.append(x[i: i + window_size])
        y_win.append(y[i + window_size])

    return np.array(x_win), np.array(y_win)


def create_sliding_window_for_early_warning(
        data: tuple,
        window_size: int,
        fire_start_idx: int,
        early_warning_window: int
) -> tuple[
    np.ndarray, np.ndarray]:

    x = data[0]
    y = data[1]
    print("Early war window creation")
    print(x.shape)
    print("Fire start index: ", fire_start_idx)
    print("Early warning window: ", early_warning_window)

    x_win, y_win = [], []
    for i in range(len(x) - window_size):
        x_win.append(x[i: i + window_size])

        t = i + window_size
        distance = fire_start_idx - t

        if 0 <= distance <= early_warning_window:
            label = 1 - (distance / early_warning_window)
        else:
            label = y[i + window_size]

        y_win.append(label)

    return np.array(x_win), np.array(y_win)