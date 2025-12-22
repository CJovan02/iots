package dto

// ReadingResponse message format received from topic "data-manager/raw-readings"
type ReadingResponse struct {
	Id          uint32  `json:"id"`
	Timestamp   int64   `json:"timestamp"`   // Raw UTC timestamp
	Temperature float64 `json:"temperature"` // Air temperature, fires raise temperature
	Humidity    float64 `json:"humidity"`    // Air humidity, very high or low can indicate fire
	Tvoc        uint32  `json:"tvoc"`        // Total Volatile Organic Compounds, high numbers indicate fire
	ECo2        uint32  `json:"e_co2"`       // CO2 equivalent concentration, indirect signal for combustion
	RawHw       uint32  `json:"raw_hw"`      // Raw molecular hydrogen, additional chemical signal
	RawEthanol  uint32  `json:"raw_ethanol"` // Raw ethanol gas, additional chemical signal
	PM25        float64 `json:"pm_25"`       // Particulate matter <2.5 µm, smoke increases concentrations
	FireAlarm   uint32  `json:"fire_alarm"`  // Ground truth, 1 if fire is present
}
