namespace Gateway.Exceptions;

public class ReadingNullResponseException : ReadingServiceException
{
    public ReadingNullResponseException() { }

    public ReadingNullResponseException(string message) : base(message) { }

    public ReadingNullResponseException(string message, Exception inner) : base(inner, message) { }
}