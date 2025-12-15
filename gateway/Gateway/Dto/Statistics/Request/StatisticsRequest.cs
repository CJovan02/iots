namespace Gateway.Dto.Statistics.Request;

public sealed record StatisticsRequest(
    DateTimeOffset StartTime,
    DateTimeOffset EndTime
);