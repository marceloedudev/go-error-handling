package exceptions

import (
	"go-error-handling/validators"
	"net/http"
)

func NewInternalServerError(message string, err error) *HttpException {
	return NewHttpException(message, http.StatusInternalServerError, validators.NewValidatorError(err))
}
