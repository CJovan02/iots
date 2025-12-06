package sensor

import "time"

// Reading represents necessary columns from the dataset for fire detection
type Reading struct {
	Id          int32     // Signed four-byte integer in the db
	Timestamp   time.Time // UTC timestamp
	Temperature float64   // Air temperature, fires raise temperature
	Humidity    float64   // Air humidity, very high or low can indicate fire
	TVOC        uint16    // Total Volatile Organic Compounds, high numbers indicate fire
	ECO2        uint16    // CO2 equivalent concentration, indirect signal for combustion
	RawHw       uint16    // Raw molecular hydrogen, additional chemical signal
	RawEthanol  uint16    // Raw ethanol gas, additional chemical signal
	PM25        float64   // Particulate matter <2.5 µm, smoke increases concentrations
	FireAlarm   uint8     // Ground truth, 1 if fire is present
}
