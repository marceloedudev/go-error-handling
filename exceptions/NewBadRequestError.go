package exceptions

import (
	"go-error-handling/validators"
	"net/http"
)

func NewBadRequestError(message string, err error) *HttpException {
	var causes []string
	if err != nil {
		causes = validators.NewValidatorError(err)
	}
	return NewHttpException(message, http.StatusBadRequest, causes)
}
