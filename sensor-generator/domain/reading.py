from dataclasses import dataclass

@dataclass
class Reading:
    timestamp: int
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
            timestamp=int(data['UTC']),
            temperature=float(data["Temperature[C]"]),
            humidity=float(data["Humidity[%]"]),
            tvoc=int(data["TVOC[ppb]"]),
            eco2=int(data["eCO2[ppm]"]),
            rawHw=int(data["Raw H2"]),
            rawEthanol=int(data["Raw Ethanol"]),
            pm25=float(data["PM2.5"]),
            fireAlarm=int(data["Fire Alarm"])
        )

    def to_dict(self) -> dict:
        return {
            "timestamp": self.timestamp,
            "temperature": self.temperature,
            "humidity": self.humidity,
            "tvoc": self.tvoc,
            "eco2": self.eco2,
            "rawHw": self.rawHw,
            "rawEthanol": self.rawEthanol,
            "pm25": self.pm25,
            "fireAlarm": self.fireAlarm
        }
