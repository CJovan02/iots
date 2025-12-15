using Gateway.Dto.Reading.Response;

namespace Gateway.Clients;

public interface IGrpcReadingClient
{
    Task<uint> CountAllAsync();
    Task<ReadingResponse> GetAsync(uint id);
}