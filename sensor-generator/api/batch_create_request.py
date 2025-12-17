from dataclasses import dataclass

from domain.reading import Reading


@dataclass
class BatchCreateRequest:
    readings: list[Reading]

    def to_dict(self) -> dict:
        return {
            "readings": [r.to_dict() for r in self.readings]
        }