package exceptions

import "net/http"

func NewNotFoundError(message string) *HttpException {
	return NewHttpException(message, http.StatusNotFound, nil)
}
