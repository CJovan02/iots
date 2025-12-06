package sensor

type Statistics struct {
	ReadingsCount    uint16
	MinTemperature   float64
	MaxTemperature   float64
	AvgTemperature   float64
	MinHumidity      float64
	MaxHumidity      float64
	AvgHumidity      float64
	SumTVOC          uint32
	FireAlarmCount   uint16
	NoFireAlarmCount uint16
}
