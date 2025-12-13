namespace Gateway.Exceptions;

public class ReadingNotFoundException : ReadingServiceException
{
    public ReadingNotFoundException()
    {
    }

    public ReadingNotFoundException(string message) : base(message)
    {
    }

    public ReadingNotFoundException(string message, Exception innerException) : base(innerException, message)
    {
    }
}