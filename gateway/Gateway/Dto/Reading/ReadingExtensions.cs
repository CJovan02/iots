using Gateway.Dto.Reading.Request;
using Gateway.Protos;

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
            Timestamp = query.Timestamp,
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

    public static BatchCreateReadingsRequest ToProto(this List<CreateReadingQuery> readings)
    {
        var request = new BatchCreateReadingsRequest();
        request.ReadingRequests.AddRange(
            readings.Select(query => query.ToProto())
        );
        return request;
    }

    public static UpdateReadingRequest ToProto(this UpdateReadingQuery query, uint id)
    {
        return new UpdateReadingRequest
        {
            Id = id,
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