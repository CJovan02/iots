namespace Gateway.Dto.Statistics.Request;

public sealed record StatisticsRequest(
    long StartTime,
    long EndTime
);