using Gateway.Dto.Reading.Response;
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

    public async Task<ReadingResponse> GetAsync(uint id)
    {
        var result = await service.GetAsync(new GetReadingRequest{Id = id});
        if (result is null)
            throw new ReadingNullResponseException();

        return ReadingResponse.From(result);
    }
}