namespace Gateway.Dto.Reading.Request;

public sealed record CreateReadingQuery(
    long Timestamp,
    double Temperature,
    double Humidity,
    int Tvoc,
    int ECo2,
    int RawHw,
    int RawEthanol,
    double Pm25,
    int FireAlarm
);