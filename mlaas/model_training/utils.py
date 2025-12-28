import sys
import numpy as np
np.set_printoptions(threshold=sys.maxsize)

def print_last_20_percent(dataset, name="dataset"):
    """
    Prints basic info and values from the last 20% of a dataset.

    :param dataset: tuple (x, y)
    :param name: name for logging
    """
    x, y = dataset
    n = len(y)

    start_idx = int(0.8 * n)  # last 20%

    print(f"\n--- {name.upper()} | LAST 80% ---")
    print(f"Total samples: {n}")
    print(f"Showing from index {start_idx} to {n}")

    print("\nY values (last 80%):")
    print(y[start_idx:])

    # Ako želiš i osnovnu statistiku
    unique, counts = np.unique(y[start_idx:], return_counts=True)
    print("\nLabel distribution:")
    for u, c in zip(unique, counts):
        print(f"  label {u}: {c}")
