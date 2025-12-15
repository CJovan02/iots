using FluentValidation;
using Gateway.Api;
using Gateway.Infrastructure;

var builder = WebApplication.CreateBuilder(args);

builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwagger();

builder.AddFluentValidationEndpointFilter();
builder.Services.AddValidatorsFromAssemblyContaining<Program>();

builder.Services.AddGrpcReadingClient();

builder.Services.AddExceptionHandlers();

var app = builder.Build();

app.UseSwagger();
app.UseSwaggerUI();

app.UseExceptionHandler();
app.UseHttpsRedirection();

app.MapReadingEndpoints();

app.Run();