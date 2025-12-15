using Gateway.Protos;

namespace Gateway.Dto.Statistics.Response;

public sealed record StatisticsResponse(
    uint ReadingsCount,
    double MinTemperature,
    double MaxTemperature,
    double AvgTemperature,
    double MinHumidity,
    double MaxHumidity,
    double AvgHumidity,
    uint SumTvoc,
    uint FireAlamCount,
    uint NoFireAlarmCount
)
{
    public static StatisticsResponse From(GetStatisticsResponse proto)
    {
        return new StatisticsResponse(
            proto.ReadingsCount,
            proto.MinTemperature,
            proto.MaxTemperature,
            proto.AvgTemperature,
            proto.MinHumidity,
            proto.MaxHumidity,
            proto.AvgHumidity,
            proto.SumTvoc,
            proto.FireAlarmCount,
            proto.NoFireAlarmCount
        );
    }
}