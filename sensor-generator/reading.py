from dataclasses import dataclass

@dataclass
class Reading:
    timestamp: str
    temperature: float
    humidity: float
    tvoc: int
    eco2: int
    rawHw: int
    rawEthanol: int
    pm25: float
    fireAlarm: int

    @staticmethod
    def from_dict(data: dict) -> 'Reading':
        return Reading(
            timestamp=data['UTC'],
            temperature=float(data["Temperature[C]"]),
            humidity=float(data["Humidity[%]"]),
            tvoc=int(data["TVOC[ppb]"]),
            eco2=int(data["eCO2[ppm]"]),
            rawHw=int(data["Raw H2"]),
            rawEthanol=int(data["Raw Ethanol"]),
            pm25=float(data["PM2.5"]),
            fireAlarm=int(data["Fire Alarm"])
        )
