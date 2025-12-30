from typing import Annotated

from pydantic import BaseModel, AfterValidator

from app.dto.reading import Reading


def has_appropriate_length(value: list[Reading]) -> list[Reading]:
    if len(value) > 40:
        raise ValueError("readings can't be more than 40 in length")
    return value


class PredictionRequest(BaseModel):
    readings: Annotated[list[Reading], AfterValidator(has_appropriate_length)]
