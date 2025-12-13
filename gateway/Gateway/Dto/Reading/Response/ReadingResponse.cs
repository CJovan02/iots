namespace Gateway.Dto.Reading.Response;

public sealed record ReadingResponse(
    uint id,
    DateTimeOffset timestamp,
    double temperature,
    double humidity,
    uint tvoc,
    uint e_co2,
    uint raw_hw,
    uint raw_ethanol,
    double pm_25,
    uint fire_alarm
);