using FluentValidation;

namespace Gateway.Dto.Statistics.Request;

public class StatisticsRequestValidator : AbstractValidator<StatisticsRequest>
{
    public StatisticsRequestValidator()
    {
        RuleFor(x => x.StartTime)
            .NotNull()
            .WithMessage("Start time cannot be null")
            .GreaterThan(0)
            .WithMessage("Start time must be greater than zero");

        RuleFor(x => x.EndTime)
            .NotNull()
            .WithMessage("End time cannot be null")
            .GreaterThan(0)
            .WithMessage("End time must be greater than zero");
    }
}