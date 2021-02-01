package exceptions

import (
	"go-error-handling/validators"
	"net/http"
)

func NewTooManyRequestsError(message string, err error) *HttpException {
	return NewHttpException(message, http.StatusTooManyRequests, validators.NewValidatorError(err))
}
