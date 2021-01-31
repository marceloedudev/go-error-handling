package exceptions

import (
	"go-error-handling/validators"
	"net/http"
)

func NewUnauthorizedError(message string, err error) *HttpException {
	var causes []string
	if err != nil {
		causes = validators.NewValidatorError(err)
	}
	return NewHttpException(message, http.StatusUnauthorized, causes)
}
