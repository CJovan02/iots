using FluentValidation;

namespace Gateway.Dto.Reading.Request;

public class UpdateReadingQueryValidator : AbstractValidator<UpdateReadingQuery>
{
    public UpdateReadingQueryValidator()
    {
        RuleFor(x => x.Id)
            .GreaterThanOrEqualTo(0)
            .WithMessage("Id must be greater than or equal to 0");

        RuleFor(x => x.ECo2)
            .GreaterThanOrEqualTo(0)
            .WithMessage("Eco2 must be greater than or equal to 0");

        RuleFor(x => x.FireAlarm)
            .GreaterThanOrEqualTo(0)
            .WithMessage("FireAlarm must be greater than or equal to 0");

        RuleFor(x => x.RawEthanol)
            .GreaterThanOrEqualTo(0)
            .WithMessage("RawEthanol must be greater than or equal to 0");

        RuleFor(x => x.RawHw)
            .GreaterThanOrEqualTo(0)
            .WithMessage("RawHw must be greater than or equal to 0");

        RuleFor(x => x.Tvoc)
            .GreaterThanOrEqualTo(0)
            .WithMessage("Tvoc must be greater than or equal to 0");
    }
}