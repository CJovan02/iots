using Gateway.Api;
using Gateway.Infrastructure;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwagger();

builder.Services.AddGrpcReadingClient();

var app = builder.Build();
// if (app.Environment.IsDevelopment())
// {
    app.UseSwagger();
    app.UseSwaggerUI();


    //app.MapOpenApi();
// }

app.UseHttpsRedirection();

app.MapReadingEndpoints();

app.Run();