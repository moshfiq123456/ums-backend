package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

// FormatValidationError converts validator errors into a friendly map
func FormatValidationError(err error) map[string]string {
	errorsMap := make(map[string]string)
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			field := strings.ToLower(e.Field())
			switch e.Tag() {
			case "required":
				errorsMap[field] = "is required"
			case "min":
				errorsMap[field] = "value is too short"
			case "max":
				errorsMap[field] = "value is too long"
			case "email":
				errorsMap[field] = "must be a valid email"
			case "e164":
				errorsMap[field] = "must be a valid phone number"
			default:
				errorsMap[field] = "invalid value"
			}
		}
	} else {
		errorsMap["error"] = err.Error()
	}
	return errorsMap
}
