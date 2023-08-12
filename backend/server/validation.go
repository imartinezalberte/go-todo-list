package server

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func ShouldBindWith(c *gin.Context, input interface{}, binds ...binding.Binding) error {
	var err error
	for _, b := range binds {
		err = c.ShouldBindWith(input, b)
		if err != nil {
			break
		}
	}

	if err, ok := err.(validator.ValidationErrors); ok {
		MapValidationErrors(input, err, "")
		return err
	}

	return err
}

func MapValidationErrors(
	input interface{},
	errs validator.ValidationErrors,
	tagKey string,
) []error {
	errors := make([]error, 0, len(errs))
	for _, err := range errs {
		fmt.Println(err)
		field, _ := reflect.TypeOf(input).Elem().FieldByName(err.Field())
		fmt.Println(field)
		// errors = append(errors, mapFieldError(err, field, tagKey))
	}

	return errors
}

func mapFieldError(fErr validator.FieldError, field reflect.StructField, tagKey string) error {
	switch fErr.Tag() {
	case "required":
		return &FieldValidationError{
			Type:  fErr.Tag(),
			Field: field.Tag.Get(tagKey),
		}

	default:
		return fErr
	}
}

type FieldValidationError struct {
	Type  string
	Field string
}

func (f *FieldValidationError) Error() string {
	return fmt.Sprintf("field %s is %s", f.Field, f.Type)
}
