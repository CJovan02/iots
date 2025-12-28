import numpy as np

def create_sliding_window(data: tuple, window_size: int) -> tuple[
    np.ndarray, np.ndarray]:
    """
    Creates sliding window for provided set
    :param data: Tuple containing x and y values
    :param window_size: The window size :)
    :return: Result tuple containing transformed model_test_results with sliding window
    """

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
) -> tuple[np.ndarray, np.ndarray]:
    """
    Creates sliding window for provided set of model_test_results, BUT, as the window that is in the calm state approaches \
    fire state, it linearly increases the label value, going from 0 to 1. So value of, ex. 0.7 means the \
    fire is very close

    :param data: Set of model_test_results containing x and y values
    :param window_size: Window size :)
    :param fire_start_idx: At what index does the fire start
    :param early_warning_window: How many readings before *fire_start_idx* to start linearly increasing the label value
    :return: Transformed set of model_test_results
    """


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