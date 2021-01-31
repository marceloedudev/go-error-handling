package validators

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func NewValidatorError(err error) []string {
	res := []string{}
	errs := err.(validator.ValidationErrors)

	for _, v := range errs {
		switch v.Tag() {
		case "required":
			{
				res = append(res, v.Field()+" is required")
			}
		case "email":
			{
				res = append(res, fmt.Sprintf("%s must be a valid email", v.Field()))
			}
		case "min":
			{
				res = append(res, fmt.Sprintf("%s must be longer than %s", v.Field(), v.Param()))
			}
		case "max":
			{
				res = append(res, fmt.Sprintf("%s cannot be longer than %s", v.Field(), v.Param()))
			}
		default:
			{
				res = append(res, fmt.Sprintf("%s is not valid", v.Field()))
			}
		}
	}

	return res
}
