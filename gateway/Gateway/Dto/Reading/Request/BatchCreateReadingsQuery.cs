namespace Gateway.Dto.Reading.Request;

public sealed record BatchCreateReadingsQuery(
    List<CreateReadingQuery> readings
);
