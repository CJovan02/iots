package domain

type ThresholdType string

const (
	ThresholdPM25        ThresholdType = "PM25"
	ThresholdTVOC        ThresholdType = "TVOC"
	ThresholdECO2        ThresholdType = "ECO2"
	ThresholdTemperature ThresholdType = "TEMPERATURE"
)

type Trigger struct {
	Type      ThresholdType
	Value     float64
	Threshold float64
}

type SmokeEvent struct {
	ReadingId uint32     `json:"reading_id"` // ReadingId in database
	Timestamp int64      `json:"timestamp"`  // Timestamp of the reading
	Triggers  []*Trigger `json:"triggers"`   // Triggers are field values that passed a certain threshold
}

func NewPm25Trigger(value float64, threshold float64) *Trigger {
	return &Trigger{
		Type:      ThresholdPM25,
		Value:     value,
		Threshold: threshold,
	}
}

func NewTvocTrigger(value float64, threshold float64) *Trigger {
	return &Trigger{
		Type:      ThresholdTVOC,
		Value:     value,
		Threshold: threshold,
	}
}

func NewEco2Trigger(value float64, threshold float64) *Trigger {
	return &Trigger{
		Type:      ThresholdECO2,
		Value:     value,
		Threshold: threshold,
	}
}

func NewTemperatureTrigger(value float64, threshold float64) *Trigger {
	return &Trigger{
		Type:      ThresholdTemperature,
		Value:     value,
		Threshold: threshold,
	}
}
