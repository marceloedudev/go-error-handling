package exceptions

import (
	"go-error-handling/validators"
	"net/http"
)

func NewUnauthorizedError(message string, err error) *HttpException {
	return NewHttpException(message, http.StatusUnauthorized, validators.NewValidatorError(err))
}
