using Gateway.Dto.Reading.Request;
using Gateway.Protos;
using Google.Protobuf.WellKnownTypes;

namespace Gateway.Dto.Reading;

public static class ReadingExtensions
{
    public static ListReadingsRequest ToProto(this ListReadingsQuery query)
    {
        return new ListReadingsRequest
        {
            PageNumber = (uint)query.PageNumber,
            PageSize = (uint)query.PageSize,
        };
    }

    public static CreateReadingRequest ToProto(this CreateReadingQuery query)
    {
        return new CreateReadingRequest
        {
            Timestamp = query.Timestamp.ToTimestamp(),
            Temperature = query.Temperature,
            Humidity = query.Humidity,
            Tvoc = (uint)query.Tvoc,
            ECo2 = (uint)query.ECo2,
            RawHw = (uint)query.RawHw,
            RawEthanol = (uint)query.RawEthanol,
            Pm25 =  query.Pm25,
            FireAlarm = (uint)query.FireAlarm,
        };
    }

    public static UpdateReadingRequest ToProto(this UpdateReadingQuery query)
    {
        return new UpdateReadingRequest
        {
            Id = (uint)query.Id,
            Temperature = query.Temperature,
            Humidity = query.Humidity,
            Tvoc = (uint)query.Tvoc,
            ECo2 = (uint)query.ECo2,
            RawHw = (uint)query.RawHw,
            RawEthanol = (uint)query.RawEthanol,
            Pm25 =  query.Pm25,
            FireAlarm = (uint)query.FireAlarm,
        };
    }
}