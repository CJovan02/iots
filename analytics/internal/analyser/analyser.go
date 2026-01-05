package analyser

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/cjovan02/iots/analytics/internal/domain"
)

// Analyser is abstraction layer for calling MLAAS REST api to analyze the reading data
type Analyser struct {
	mlaasUrl   string
	windowSize uint // windowSize - when this amount of readings get collected for one type of key (deviceId), send it to MLAAS for analysing
	httpClient *http.Client

	windows map[string][]domain.Reading // windows - groups readings by *deviceId*. This way, only when _windowSize_ amount of readings get collected for that *readingId* it will send it to MLAAS to analyse it
}

func NewAnalyser(mlaasUrl string, windowSize uint) *Analyser {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	return &Analyser{
		mlaasUrl:   mlaasUrl,
		windowSize: windowSize,
		httpClient: &client,
		windows:    make(map[string][]domain.Reading),
	}
}

// AddReading - Adds reading to the analyser.
// When enough readings get collected for one type of ReadingId, it will call MLAAS to analyse that data.
// nil, false, nil is returned when you add the reading but there are not enough readings for that deviceId
// response, true, nil is returned when there are enough readings and data is analysed
func (a *Analyser) AddReading(reading domain.Reading) (*PredictionResponse, bool, error) {
	key := reading.DeviceId

	// Initialize the array if it doesn't exist
	if a.windows[key] == nil {
		a.windows[key] = make([]domain.Reading, 0)
	}

	window := a.windows[key]
	window = append(window, reading)

	winLen := len(window)
	if winLen < int(a.windowSize) {
		a.windows[key] = window
		log.Printf(
			"successfully added reading to the group %s. Window now has %d readings\n", reading.DeviceId, winLen)
		return nil, false, nil
	}

	log.Printf("group %s, has enough data to analyse\n", reading.DeviceId)

	// rest window for that key. With this method I think we avoid reallocating the memory
	a.windows[key] = window[:0]
	request := FromDomainReadings(window)
	resp, err := a.predict(request)
	if err != nil {
		return nil, false, err
	}

	return resp, true, nil
}

func (a *Analyser) predict(request *PredictionRequest) (*PredictionResponse, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	url := a.mlaasUrl + "/predict"
	res, err := a.httpClient.Post(url, "application/json", bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("mlaas responded with status %d. Body: \n%s", res.StatusCode, string(body))
	}

	var resp PredictionResponse
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &resp, nil
}
