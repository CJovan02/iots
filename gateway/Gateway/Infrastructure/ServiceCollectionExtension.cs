using Gateway.Clients;
using Gateway.Protos;
using Grpc.Net.Client;
using Microsoft.OpenApi.Models;

namespace Gateway.Infrastructure;

public static class ServiceCollectionExtension
{
    public static IServiceCollection AddGrpcReadingClient(this IServiceCollection services)
    {
        var address = "http://localhost:8080";

        services.AddSingleton(sp =>
        {
            var channel = GrpcChannel.ForAddress(address);
            return new Readings.ReadingsClient(channel);
        });

        services.AddScoped<IGrpcReadingClient, GrpcReadingClient>();

        return services;
    }

    public static IServiceCollection AddSwagger(this IServiceCollection services)
    {
        return services.AddSwaggerGen(options =>
        {
            options.SwaggerDoc("v1", new OpenApiInfo { Title = "Gateway Api", Version = "v1" });
        });
    }
}