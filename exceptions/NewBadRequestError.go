package exceptions

import (
	"go-error-handling/validators"
	"net/http"
)

func NewBadRequestError(message string, err error) *HttpException {
	return NewHttpException(message, http.StatusBadRequest, validators.NewValidatorError(err))
}
