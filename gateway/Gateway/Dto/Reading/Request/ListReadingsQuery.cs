namespace Gateway.Dto.Reading.Request;

public sealed record ListReadingsQuery(
    int PageNumber,
    int PageSize
);