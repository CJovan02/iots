from tensorflow.keras.models import Sequential
from keras.src.losses import mean_absolute_error

from matplotlib import pyplot as plt

import pandas as pd

import numpy as np


def test_predictions(
        model: Sequential,
        early_warning_test: tuple[np.ndarray, np.ndarray],
        fire_test: tuple[np.ndarray, np.ndarray],
        calm_test: tuple[np.ndarray, np.ndarray],
) -> None:
    """
    Preforms predictions on test sets provided, shows Mean Absolute Error and plots predicted results along side real results

    :param model: Trained model
    :param early_warning_test: Set before fire 2 starts
    :param calm_test: Set during the fire
    :param fire_test: Set during calm period
    :return:
    """

    # Early fire 2
    early_war_pred = model.predict(early_warning_test[0]).flatten()

    print("Early fire 2 warning test")
    print(mean_absolute_error(early_warning_test[1], early_war_pred))

    early_war_results = pd.DataFrame(data={"Train Predictions": early_war_pred, "Actual Values": early_warning_test[1]})
    plt.plot(early_war_results["Train Predictions"])
    plt.plot(early_war_results["Actual Values"])
    plt.show()

    # Fire test
    fire_test_pred = model.predict(fire_test[0]).flatten()

    print("Fire test")
    print(mean_absolute_error(fire_test[1], fire_test_pred))
    fire_test_results = pd.DataFrame(data={"Train Predictions": fire_test_pred, "Actual Values": fire_test[1]})
    plt.plot(fire_test_results["Train Predictions"])
    plt.plot(fire_test_results["Actual Values"])
    plt.show()

    # Calm test
    calm_test_pred = model.predict(calm_test[0]).flatten()

    print("Calm test")
    print(mean_absolute_error(calm_test[1], calm_test_pred))
    calm_test_results = pd.DataFrame(data={"Train Predictions": calm_test_pred, "Actual Values": calm_test[1]})
    plt.plot(calm_test_results["Train Predictions"])
    plt.plot(calm_test_results["Actual Values"])
    plt.show()
