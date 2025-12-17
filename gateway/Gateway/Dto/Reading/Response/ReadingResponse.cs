using Gateway.Protos;

namespace Gateway.Dto.Reading.Response;

public sealed record ReadingResponse(
    uint Id,
    DateTime Timestamp,
    double Temperature,
    double Humidity,
    uint Tvoc,
    uint ECo2,
    uint RawHw,
    uint RawEthanol,
    double Pm25,
    uint FireAlarm
)
{
    public static ReadingResponse From(GetReadingResponse proto)
    {
        return new ReadingResponse
        (
            proto.Id,
            DateTimeOffset.FromUnixTimeSeconds(proto.Timestamp).UtcDateTime,
            proto.Temperature,
            proto.Humidity,
            proto.Tvoc,
            proto.ECo2,
            proto.RawHw,
            proto.RawEthanol,
            proto.Pm25,
            proto.FireAlarm
        );
    }
}