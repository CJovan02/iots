using FluentValidation;

namespace Gateway.Dto.Reading.Request;

public class ListReadingsQueryValidator : AbstractValidator<ListReadingsQuery>
{
    public ListReadingsQueryValidator()
    {
        RuleFor(x => x.PageNumber)
            .NotEmpty()
            .WithMessage("PageNumber must be greater than zero")
            .GreaterThan(0)
            .WithMessage("PageNumber must be greater than zero");

        RuleFor(x => x.PageSize)
            .NotEmpty()
            .WithMessage("PageSize must be greater than zero")
            .GreaterThan(0)
            .WithMessage("PageSize must be greater than zero");
    }
}