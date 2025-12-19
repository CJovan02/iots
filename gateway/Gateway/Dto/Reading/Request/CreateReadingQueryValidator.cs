using FluentValidation;

namespace Gateway.Dto.Reading.Request;

public class CreateReadingQueryValidator : AbstractValidator<CreateReadingQuery>
{
    public CreateReadingQueryValidator()
    {
        RuleFor(x => x.Timestamp)
            .NotEmpty()
            .WithMessage("Timestamp must not be empty")
            .GreaterThan(0)
            .WithMessage("Timestamp must be greater than 0");

        RuleFor(x => x.Temperature)
            .NotNull()
            .WithMessage("Temperature must not be null");

        RuleFor(x => x.Humidity)
            .NotEmpty()
            .WithMessage("Humidity must not be empty")
            .GreaterThan(0)
            .WithMessage("Humidity must be greater than 0");

        RuleFor(x => x.Tvoc)
            .NotNull()
            .WithMessage("Tvoc must not be null")
            .GreaterThanOrEqualTo(0)
            .WithMessage("Tvoc must be greater than or equal to 0");

        RuleFor(x => x.ECo2)
            .NotNull()
            .WithMessage("ECo2 must not be null")
            .GreaterThanOrEqualTo(0)
            .WithMessage("Eco2 must be greater than or equal to 0");

        RuleFor(x => x.RawEthanol)
            .NotNull()
            .WithMessage("RawEthanol must not be null")
            .GreaterThanOrEqualTo(0)
            .WithMessage("RawEthanol must be greater than or equal to 0");

        RuleFor(x => x.RawHw)
            .NotNull()
            .WithMessage("RawHw must not be null")
            .GreaterThanOrEqualTo(0)
            .WithMessage("RawHw must be greater than or equal to 0");

        RuleFor(x => x.Pm25)
            .NotNull()
            .WithMessage("Pm25 must not be null")
            .GreaterThanOrEqualTo(0)
            .WithMessage("Pm25 must be greater than or equal to 0");

        RuleFor(x => x.FireAlarm)
            .NotNull()
            .WithMessage("FireAlarm must not be null")
            .GreaterThanOrEqualTo(0)
            .WithMessage("FireAlarm must be greater than or equal to 0");
    }
}