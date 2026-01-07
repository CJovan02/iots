package nats

type PredictionMessage struct {
	Prediction    float64  `json:"prediction"`     // Float number between 0 and 1. The closer the number is to the 1, the greater the risk of fire in the near future
	DeviceId      string   `json:"device_id"`      // Device ID that recorded the list of readings used for prediction
	ReadingsCount uint32   `json:"readings_count"` // Number of readings used for the prediction
	ReadingIds    []uint32 `json:"reading_ids"`    // Ids of the readings in database used for prediction
	TimeInterval  int32    `json:"time_interval"`  // Time interval between last and first reading in seconds
}
