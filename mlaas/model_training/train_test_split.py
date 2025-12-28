import sys
import numpy as np

np.set_printoptions(threshold=sys.maxsize)


# Yes, it's a big function :)
def create_train_test_split(x, y, early_warning_interval: int, future_prediction: int = 0) -> tuple[
    tuple[np.ndarray, np.ndarray], tuple[np.ndarray, np.ndarray], tuple[np.ndarray, np.ndarray], tuple[
        np.ndarray, np.ndarray]]:
    """
    We want AI model to be able to predict fire before it happens. In dataset, we only have 2 fire events that last for
    quite a while. It would be ideal if we had dataset that has a lot of fire events.
    Since we don't have that, we improvise:

    - Right before the fire starts, we grab *early_warning_interval* amount of readings and assign it to "early_warning_test"
    - Around 15% of readings at the end of fire events and assign it "fire_test"
    - Around 10% of readings at the end of a dataset that represent calm state and assign it to "calm_test"
    - The rest of the model_test_results is used for training.

    We do this process for both fire events inside dataset.

    The "early_warning_test" is the most important, but since we don't have enough fire events to make it big enough.
    we create 2 more tests that have way more model_test_results for testing.

    :param x: Features
    :param y: Label
    :param early_warning_interval: The amount of readings before the fire to put into "early_warning_test"
    :param future_prediction: How many readings model should predict in the future. When splitting model_test_results in test set we need to account for this metric.
    :return:
    """

    if early_warning_interval > 3000 or early_warning_interval < 0:
        raise AttributeError("early_warning_interval must be between 0 and 3000")

    fire_1_start = 3179  # inclusive
    fire_1_end = 24995  # exclusive
    fire_2_start = 28173  # inclusive
    fire_2_end = 51143  # exclusive
    # Percent amount of fire event to grab and put inside test set
    fire_amount = 0.15

    fire_1_length = fire_1_end - fire_1_start
    fire_2_length = fire_2_end - fire_2_start

    print("Fire 1 and fire 2 lengths: ")
    print(fire_1_length, fire_2_length)
    print()

    # We start from the beginning, grabbing everything from the start and up to 85% of entire fire 1
    train_set = (
        x[:fire_1_end - int(fire_1_length * fire_amount)],
        y[:fire_1_end - int(fire_1_length * fire_amount)]
    )

    print("Train set, first part and 85% of fire 1: ")
    print(train_set[0].shape, train_set[1].shape)
    print()

    # Then we grab the remaining of the fire 1 event for test set
    fire_test = (
        x[fire_1_end - int(fire_1_length * fire_amount):fire_1_end],
        y[fire_1_end - int(fire_1_length * fire_amount):fire_1_end],
    )

    print("Fire test for fire 1:")
    print(fire_test[0].shape, fire_test[1].shape)
    print()

    # We are done with first fire event
    # Now we start all over again for fire 2 and combine with previous results
    train_set_before_fire_2_x = x[fire_1_end:fire_2_start - early_warning_interval]
    train_set_before_fire_2_y = y[fire_1_end:fire_2_start - early_warning_interval]

    print("Train set before fire 2: ")
    print(train_set_before_fire_2_x.shape, train_set_before_fire_2_y.shape)
    print()

    train_set = (
        np.concatenate([train_set[0], train_set_before_fire_2_x], axis=0),
        np.concatenate([train_set[1], train_set_before_fire_2_y], axis=0)
    )

    # We grab the early warning interval before fire 2
    early_warning_test_2 = (
        x[fire_2_start - early_warning_interval:fire_2_start + future_prediction],
        y[fire_2_start - early_warning_interval:fire_2_start + future_prediction]
    )

    print("Early warning test for fire 2: ")
    print(early_warning_test_2[0].shape, early_warning_test_2[1].shape)
    # print(early_war_fire_2_y)
    print()

    # Then we grab first 85% of the fire 2 event for training
    fire_2_train_x = x[fire_2_start + future_prediction:fire_2_end - int(fire_2_length * fire_amount)]
    fire_2_train_y = y[fire_2_start + future_prediction:fire_2_end - int(fire_2_length * fire_amount)]

    print("Fire 2 for train: ")
    print(fire_2_train_x.shape, fire_2_train_y.shape)
    print()

    train_set = (
        np.concatenate([train_set[0], fire_2_train_x], axis=0),
        np.concatenate([train_set[1], fire_2_train_y], axis=0)
    )

    # Last 15% of fire 2 goes to testing...
    fire_2_test_x = x[fire_2_end - int(fire_2_length * fire_amount):fire_2_end]
    fire_2_test_y = y[fire_2_end - int(fire_2_length * fire_amount):fire_2_end]

    print("Fire 2 for test: ")
    print(fire_2_test_x.shape, fire_2_test_y.shape)
    print()

    fire_test = (
        np.concatenate([fire_test[0], fire_2_test_x], axis=0),
        np.concatenate([fire_test[1], fire_2_test_y], axis=0)
    )

    # We get another 8000 readings after fire 2 for training
    train_set = (
        np.concatenate([train_set[0], x[fire_2_end:fire_2_end + 5000]], axis=0),
        np.concatenate([train_set[1], y[fire_2_end:fire_2_end + 5000]], axis=0)
    )

    # The rest goes to clam_test
    calm_test = (
        x[fire_2_end + 5000:],
        y[fire_2_end + 5000:]
    )

    print("Train set: ")
    print(train_set[0].shape, train_set[1].shape)
    print()

    print("Early warning test 2: ")
    print(early_warning_test_2[0].shape, early_warning_test_2[1].shape)
    # print(early_warning_test[1])
    print()

    print("Fire test")
    print(fire_test[0].shape, fire_test[1].shape)
    print()

    print("Calm test: ")
    print(calm_test[0].shape, calm_test[1].shape)
    print()

    return train_set, early_warning_test_2, fire_test, calm_test
