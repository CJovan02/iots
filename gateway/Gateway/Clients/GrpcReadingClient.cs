using Gateway.Exceptions;
using Gateway.Protos;
using Google.Protobuf.WellKnownTypes;
using Grpc.Core;

namespace Gateway.Clients;

public class GrpcReadingClient(Readings.ReadingsClient service) : IGrpcReadingClient
{
    public async Task<uint> CountAll()
    {
        try
        {
            var result = await service.CountAllAsync(new Empty());
            if (result is null)
                throw new ReadingNullResponseException();

            return result.Count;
        }
        catch (RpcException ex)
        {
            throw new ReadingServiceException(ex);
        }
    }
}