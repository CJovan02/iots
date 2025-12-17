from dataclasses import dataclass

from reading import Reading


@dataclass
class BatchCreateRequest:
    readings: list[Reading]