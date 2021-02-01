package exceptions

import (
	"go-error-handling/validators"
	"net/http"
)

func NewNotAcceptableError(message string, err error) *HttpException {
	return NewHttpException(message, http.StatusNotAcceptable, validators.NewValidatorError(err))
}
