using Gateway.Clients;

namespace Gateway.Api;

public static class ReadingEndpoints
{
    public static void MapReadingEndpoints(this WebApplication app)
    {
        app.MapGet("readings/count", async (IGrpcReadingClient client) =>
        {
            var count = await client.CountAllAsync();
            return Results.Ok(count);
        })
        .WithName("GetReadingsCount")
        .WithTags("Readings");
    }
}