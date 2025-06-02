package types

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type XValidator struct {
	Validator *validator.Validate
}

type ValidationErrorResponse struct {
	FailedField string
	Tag         string
	Value       any
}

func (v *XValidator) Validate(data any) []ValidationErrorResponse {
	errors := make([]ValidationErrorResponse, 0)

	errs := v.Validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			validationError := &ValidationErrorResponse{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Value(),
			}

			errors = append(errors, *validationError)
		}
	}
	return errors
}

func ParseRequest[T any](v *XValidator, c *fiber.Ctx, r *T) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}

	if errs := v.Validate(r); len(errs) > 0 {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		return errors.New(strings.Join(errMsgs, " and "))
	}

	return nil
}
