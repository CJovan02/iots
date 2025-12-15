using Gateway.Dto.Reading;
using Gateway.Dto.Reading.Request;
using Gateway.Dto.Reading.Response;
using Gateway.Dto.Statistics;
using Gateway.Dto.Statistics.Request;
using Gateway.Dto.Statistics.Response;
using Gateway.Exceptions;
using Gateway.Protos;
using Google.Protobuf.WellKnownTypes;

namespace Gateway.Clients;

public class GrpcReadingClient(Readings.ReadingsClient service) : IGrpcReadingClient
{
    public async Task<uint> CountAllAsync()
    {
        var result = await service.CountAllAsync(new Empty());
        if (result is null)
            throw new ReadingNullResponseException();

        return result.Count;
    }

    public async Task<List<ReadingResponse>> ListAsync(ListReadingsQuery query)
    {
        var result = await service.ListAsync(query.ToProto());
        if (result is null)
            throw new ReadingNullResponseException();

        return result.Readings
            .Select(ReadingResponse.From)
            .ToList();
    }

    public async Task<ReadingResponse> GetAsync(uint id)
    {
        var result = await service.GetAsync(new GetReadingRequest { Id = id });
        if (result is null)
            throw new ReadingNullResponseException();

        return ReadingResponse.From(result);
    }

    public async Task<StatisticsResponse> StatisticsAsync(StatisticsRequest request)
    {
        var result = await service.StatisticsAsync(request.ToProto());
        if (result is null)
            throw new ReadingNullResponseException();

        return StatisticsResponse.From(result);
    }

    public async Task<uint> CreateAsync(CreateReadingQuery request)
    {
        var result = await service.CreateAsync(request.ToProto());
        if (result is null)
            throw new ReadingNullResponseException();

        return result.Id;
    }

    public async Task UpdateAsync(uint id, UpdateReadingQuery request)
    {
        var result = await service.UpdateAsync(request.ToProto(id));
        if (result is null)
            throw new ReadingNullResponseException();
    }

    public async Task DeleteAsync(uint id)
    {
        var result = await service.DeleteAsync(new DeleteReadingRequest { Id = id });
        if (result is null)
            throw new ReadingNullResponseException();
    }
}