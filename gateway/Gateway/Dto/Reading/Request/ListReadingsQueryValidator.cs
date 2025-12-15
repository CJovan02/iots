using FluentValidation;

namespace Gateway.Dto.Reading.Request;

public class ListReadingsQueryValidator : AbstractValidator<ListReadingsQuery>
{
    public ListReadingsQueryValidator()
    {
        RuleFor(x => x.PageNumber)
            .GreaterThan(0)
            .WithMessage("PageNumber must be greater than zero");

        RuleFor(x => x.PageSize)
            .GreaterThan(0)
            .WithMessage("PageSize must be greater than zero");
    }
}