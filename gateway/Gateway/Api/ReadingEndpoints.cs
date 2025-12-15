using Gateway.Clients;
using Gateway.Dto.Reading.Request;
using Gateway.Dto.Statistics.Request;
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

        app.MapGet("readings", async ([AsParameters] ListReadingsQuery query, IGrpcReadingClient client) =>
            {
                var readings = await client.ListAsync(query);
                return Results.Ok(readings);
            })
            .AddFluentValidationFilter()
            .WithName("ListReadings")
            .WithTags("Readings");

        app.MapGet("readings/statistics", async ([AsParameters] StatisticsRequest request, IGrpcReadingClient client) =>
            {
                var statistics = await client.StatisticsAsync(request);
                return Results.Ok(statistics);
            })
            .AddFluentValidationFilter()
            .WithName("GetReadingsStatistics")
            .WithTags("Readings");

        app.MapGet("readings/{id:int}", async (int id, IGrpcReadingClient client) =>
            {
                if (id <= 0)
                    return Results.BadRequest(new ProblemDetails
                    {
                        Title = "Invalid Argument",
                        Status = 400,
                        Detail = "Id must be greater than 0",
                    });

                var reading = await client.GetAsync((uint)id);
                return Results.Ok(reading);
            })
            .WithName("GetReadingById")
            .WithTags("Readings");

        app.MapPost("readings", async ([FromBody] CreateReadingQuery query, IGrpcReadingClient client) =>
            {
                var id = await client.CreateAsync(query);
                return Results.Ok(id);
            })
            .AddFluentValidationFilter()
            .WithName("CreateReading")
            .WithTags("Readings");

        app.MapPut("readings/{id:int}", async (int id, [FromBody] UpdateReadingQuery query, IGrpcReadingClient client) =>
            {
                if (id <= 0)
                    return Results.BadRequest(new ProblemDetails
                    {
                        Title = "Invalid Argument",
                        Status = 400,
                        Detail = "Id must be greater than 0",
                    });

                await client.UpdateAsync((uint)id, query);
                return Results.Ok();
            })
            .AddFluentValidationFilter()
            .WithName("UpdateReading")
            .WithTags("Readings");

        app.MapDelete("readings/{id:int}", async (int id, IGrpcReadingClient client) =>
            {
                if (id <= 0)
                    return Results.BadRequest(new ProblemDetails
                    {
                        Title = "Invalid Argument",
                        Status = 400,
                        Detail = "Id must be greater than 0",
                    });

                await client.DeleteAsync((uint)id);
                return Results.Ok();
            })
            .WithName("DeleteReading")
            .WithTags("Readings");
    }
}