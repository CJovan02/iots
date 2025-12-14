using Grpc.Core;
using Microsoft.AspNetCore.Diagnostics;
using Microsoft.AspNetCore.Mvc;

namespace Gateway.ExceptionHandlers;

internal sealed class RpcExceptionHandler(ILogger<RpcExceptionHandler> logger) : IExceptionHandler
{
    private readonly ILogger<RpcExceptionHandler> _logger = logger;

    public async ValueTask<bool> TryHandleAsync(
        HttpContext httpContext,
        Exception exception,
        CancellationToken cancellationToken)
    {
        if (exception is not RpcException rpcException)
        {
            return false;
        }

        _logger.LogError(
            rpcException,
            "RPC exception occured: {Message}",
            rpcException.Message);

        var problemDetails = new ProblemDetails
        {
            Detail = rpcException.Status.Detail,
        };

        switch (rpcException.StatusCode)
        {
            case StatusCode.InvalidArgument:
                problemDetails.Title = "Invalid argument";
                problemDetails.Status = 400;
                break;
            case StatusCode.NotFound:
                problemDetails.Title = "Not found";
                problemDetails.Status = 404;
                break;
            case StatusCode.Unavailable:
                problemDetails.Title = "Service unavailable";
                problemDetails.Status = 503;
                break;
            default:
                problemDetails.Title = "Internal server error";
                problemDetails.Status = 500;
                break;
        }

        httpContext.Response.StatusCode = problemDetails.Status.Value;

        await httpContext.Response
            .WriteAsJsonAsync(problemDetails, cancellationToken);

        return true;
    }
}