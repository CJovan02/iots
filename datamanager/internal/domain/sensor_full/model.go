package sensor_full

// Reading - Model from the dataset, it contains all the columns from .csv file
type Reading struct {
	Index       uint16  // Dataset has 62.6k rows
	Timestamp   int64   // Raw Timestamp in UTC
	Temperature float64 // Air temperature
	Humidity    float64 // Air humidity
	TVOC        uint16  // Otal Volatile Organic Compounds; measured in parts per billion
	eCO2        uint16  // Co2 equivalent concentration; calculated from different values like TVCO
	RawHw       uint16  // Raw molecular hydrogen; not compensated (Bias, temperature, etc.)
	RawEthanol  uint16  // Raw ethanol gas
	Pressure    float64 // Air pressure in hPa
	PM10        float64 // Particulate matter size < 1.0 µm (PM1.0). 1.0 µm < 2.5 µm (PM2.5)
	PM25        float64 // Particulate matter size < 1.0 µm (PM1.0). 1.0 µm < 2.5 µm (PM2.5)
	NC05        float64 // Number concentration of particulate matter. This differs from PM because NC gives the actual number of particles in the air. The raw NC is
	NC10        float64 // Number concentration of particulate matter. This differs from PM because NC gives the actual number of particles in
	NC25        float64 // Number concentration of particulate matter. This differs from PM because NC gives the actual number of particles in
	CNT         uint16  // Sample counter
	FireAlarm   uint8   // Ground truth is "1" if a fire is there
}
