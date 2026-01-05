package analyser

import (
	"log"

	"github.com/cjovan02/iots/analytics/internal/domain"
)

// Analyser is abstraction layer for calling MLAAS REST api to analyze the reading data
type Analyser struct {
	MlaasUrl string
}

func NewAnalyser(mlaasUrl string) *Analyser {
	return &Analyser{MlaasUrl: mlaasUrl}
}

func (a *Analyser) Predict(reading domain.Reading) {
	request := FromDomainReading(reading)
	log.Println(request)
}
