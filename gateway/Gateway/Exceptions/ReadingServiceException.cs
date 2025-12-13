namespace Gateway.Exceptions;

public class ReadingServiceException : Exception
{
    public ReadingServiceException()
    {
    }

    public ReadingServiceException(string message = "gRPC call failed") : base(message)
    {
    }

    public ReadingServiceException(Exception innerException, string message = "gRPC call failed")
        : base(message, innerException)
    {
    }
}