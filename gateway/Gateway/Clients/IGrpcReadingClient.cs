using Gateway.Protos;

namespace Gateway.Clients;

public interface IGrpcReadingClient
{
    Task<uint> CountAll();
}