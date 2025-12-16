using FluentValidation;

namespace Gateway.Dto.Reading.Request;

public class BatchCreateReadingsQueryValidator : AbstractValidator<BatchCreateReadingsQuery>
{
    public BatchCreateReadingsQueryValidator()
    {
        RuleFor(x => x.readings)
            .NotEmpty()
            .WithMessage("Please specify at least one reading");

        RuleForEach(x => x.readings)
            .SetValidator(new CreateReadingQueryValidator());
    }
}