package sensor

// Reading represents necessary columns from the dataset for fire detection
type Reading struct {
	Id          uint32
	Timestamp   int64   // Raw UTC timestamp
	Temperature float64 // Air temperature, fires raise temperature
	Humidity    float64 // Air humidity, very high or low can indicate fire
	Tvoc        uint32  // Total Volatile Organic Compounds, high numbers indicate fire
	ECo2        uint32  // CO2 equivalent concentration, indirect signal for combustion
	RawHw       uint32  // Raw molecular hydrogen, additional chemical signal
	RawEthanol  uint32  // Raw ethanol gas, additional chemical signal
	PM25        float64 // Particulate matter <2.5 µm, smoke increases concentrations
	FireAlarm   uint32  // Ground truth, 1 if fire is present
}
