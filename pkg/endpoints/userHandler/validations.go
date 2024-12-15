package userHandler

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

type UserValidator struct {
	validator *validator.Validate
}

func NewValidator() UserValidator {

	v := validator.New()

	v.RegisterValidation("pan", func(fl validator.FieldLevel) bool {
		panRegex := regexp.MustCompile(`^[A-Z]{5}[0-9]{4}[A-Z]$`)
		return panRegex.MatchString(fl.Field().String())
	})

	// Custom validation for mobile number
	v.RegisterValidation("mobile", func(fl validator.FieldLevel) bool {
		mobileRegex := regexp.MustCompile(`^[0-9]{10}$`)
		return mobileRegex.MatchString(fl.Field().String())
	})

	return UserValidator{validator: v}

}
