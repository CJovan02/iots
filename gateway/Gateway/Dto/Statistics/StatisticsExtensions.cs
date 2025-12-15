using Gateway.Dto.Statistics.Request;
using Gateway.Protos;
using Google.Protobuf.WellKnownTypes;

namespace Gateway.Dto.Statistics;

public static class StatisticsExtensions
{
    public static GetStatisticsRequest ToProto(this StatisticsRequest request)
    {
        return new GetStatisticsRequest
        {
            StartTime = request.StartTime.ToTimestamp(),
            EndTime = request.EndTime.ToTimestamp(),
        };
    }
}