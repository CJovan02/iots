using Gateway.Dto.Reading.Request;
using Gateway.Dto.Reading.Response;
using Gateway.Dto.Statistics.Request;
using Gateway.Dto.Statistics.Response;

namespace Gateway.Clients;

public interface IGrpcReadingClient
{
    Task<uint> CountAllAsync();
    Task<List<ReadingResponse>> ListAsync(ListReadingsQuery query);
    Task<ReadingResponse> GetAsync(uint id);
    Task<StatisticsResponse> StatisticsAsync(StatisticsRequest request);
    Task<uint> CreateAsync(CreateReadingQuery request);
    Task UpdateAsync(UpdateReadingQuery request);
    Task DeleteAsync(uint id);
}