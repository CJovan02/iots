from typing import Annotated

from pydantic import AfterValidator, BaseModel


def is_positive(value: int) -> int:
    if value < 0:
        raise ValueError("value must be positive")
    return value


PositiveInt = Annotated[int, AfterValidator(is_positive)]


class Reading(BaseModel):
    temperature: float
    humidity: PositiveInt
    tvoc: PositiveInt
    e_co2: PositiveInt
    raw_hw: PositiveInt
    raw_ethanol: PositiveInt
    pm_25: float
