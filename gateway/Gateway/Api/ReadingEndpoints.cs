using Gateway.Clients;
using Microsoft.AspNetCore.Mvc;

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

        app.MapGet("readings/{id:int}", async (int id, IGrpcReadingClient client) =>
            {
                if (id <= 0)
                    return Results.BadRequest(new ProblemDetails
                    {
                        Title = "Bad Request",
                        Status = 400,
                        Detail = "Id must be greater than 0",
                    });

                var reading = await client.GetAsync((uint)id);
                return Results.Ok(reading);
            })
            .AddFluentValidationFilter()
            .WithName("GetReadingById")
            .WithTags("Readings");
    }
}