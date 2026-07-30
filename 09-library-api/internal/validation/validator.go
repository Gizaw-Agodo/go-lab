package validation

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func Validate(v any) error {
	if err := validate.Struct(v); err != nil {
		validationErrors := err.(validator.ValidationErrors)

		var messages []string

		for _, fieldErr := range validationErrors {
			switch fieldErr.Tag() {
			case "required":
				messages = append(messages,
					fmt.Sprintf("%s is required", strings.ToLower(fieldErr.Field())),
				)

			case "min":
				messages = append(messages,
					fmt.Sprintf("%s must be at least %s characters",
						strings.ToLower(fieldErr.Field()),
						fieldErr.Param(),
					),
				)

			case "max":
				messages = append(messages,
					fmt.Sprintf("%s must be at most %s characters",
						strings.ToLower(fieldErr.Field()),
						fieldErr.Param(),
					),
				)

			default:
				messages = append(messages, fieldErr.Error())
			}
		}

		return fmt.Errorf("%s", strings.Join(messages, ", "))
	}

	return nil
}