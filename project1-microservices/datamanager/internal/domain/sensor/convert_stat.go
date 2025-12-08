package sensor

import "github.com/CJovan02/iots/project1-microservices/datamanager/protogen/golang/sensorpg"

func (stat *Statistics) ToProto() *sensorpg.GetStatisticsResponse {
	return &sensorpg.GetStatisticsResponse{
		ReadingsCount:    stat.ReadingsCount,
		MinTemperature:   stat.MinTemperature,
		MaxTemperature:   stat.MaxTemperature,
		AvgTemperature:   stat.AvgTemperature,
		MinHumidity:      stat.MinHumidity,
		MaxHumidity:      stat.MaxHumidity,
		AvgHumidity:      stat.AvgHumidity,
		SumTvoc:          stat.SumTVOC,
		FireAlarmCount:   stat.FireAlarmCount,
		NoFireAlarmCount: stat.NoFireAlarmCount,
	}
}
