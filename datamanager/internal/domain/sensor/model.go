package sensor

// Reading represents necessary columns from the dataset for fire detection
type Reading struct {
	Id          uint32  `json:"id"`
	DeviceId    string  `json:"device_id"`   // DeviceId is the id of the device that recorded this reading
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

func NewReading(
	id uint32,
	deviceId string,
	timestamp int64,
	temperature float64,
	humidity float64,
	tvoc uint32,
	eCo2 uint32,
	rawHw uint32,
	rawEthanol uint32,
	pm25 float64,
	fireAlarm uint32,
) *Reading {
	if deviceId == "" {
		deviceId = "device-1" // Since dataset only uses one device, we hardcode this. This is useful for other microservices.
	}

	return &Reading{
		Id:          id,
		DeviceId:    deviceId,
		Timestamp:   timestamp,
		Temperature: temperature,
		Humidity:    humidity,
		Tvoc:        tvoc,
		ECo2:        eCo2,
		RawHw:       rawHw,
		RawEthanol:  rawEthanol,
		PM25:        pm25,
		FireAlarm:   fireAlarm,
	}
}
