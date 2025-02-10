package validate

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

func registerValidations(v *validator.Validate) {
	// * Do remember to register error message as well in message.go

	v.RegisterValidation("phone", validatePhone)
	v.RegisterValidation("sgphone", validateSGPhone)
}

// regexes
var (
	phoneRegex = regexp.MustCompile(`^\+?[1-9][0-9]{7,14}$`)
	sgPhone    = regexp.MustCompile(`^(\+?65)?[89]\d{7}$`)
)

func validatePhone(fl validator.FieldLevel) bool {
	return phoneRegex.MatchString(strings.ReplaceAll(fl.Field().String(), " ", ""))
}

// validateSGPhone validates Singapore phone number
func validateSGPhone(fl validator.FieldLevel) bool {
	return sgPhone.MatchString(strings.ReplaceAll(fl.Field().String(), " ", ""))
}
