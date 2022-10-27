package helper

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

type errorValidation struct {
	Errors map[string]string `json:"errors"`
}

func ValidateRequest(verr validator.ValidationErrors) errorValidation {
	errs := make(map[string]string)
	for _, f := range verr {
		regex := regexp.MustCompile("(.)([A-Z])")
		field := strings.ToLower(regex.ReplaceAllString(f.StructField(), "${1}_$2"))
		err := translate(field, f.ActualTag(), f.Param())
		errs[field] = err
	}

	return errorValidation{errs}
}

func translate(field, err, param string) string {
	if err == "required" {
		return fmt.Sprint(field + " is required")
	}

	if err == "min" {
		if field == "age" {
			return fmt.Sprintf("%s cannot be lower than %s", field, param)
		}

		return fmt.Sprintf("%s has to have a minimum length of %s characters", field, param)
	}

	return fmt.Sprint(field + " is invalid")
}
