package analyser

import "github.com/cjovan02/iots/analytics/internal/domain"

type Reading struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Tvoc        uint32  `json:"tvoc"`
	ECo2        uint32  `json:"e_co2"`
	RawHw       uint32  `json:"raw_hw"`
	RawEthanol  uint32  `json:"raw_ethanol"`
	PM25        float64 `json:"pm_25"`
}

func FromDomainReading(r domain.Reading) *Reading {
	return &Reading{
		Temperature: r.Temperature,
		Humidity:    r.Humidity,
		Tvoc:        r.Tvoc,
		ECo2:        r.ECo2,
		RawHw:       r.RawHw,
		RawEthanol:  r.RawEthanol,
		PM25:        r.PM25,
	}
}
