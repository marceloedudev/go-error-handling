package exceptions

import "time"

type HttpException struct {
	Message   string    `json:"message"`
	Status    int       `json:"status"`
	Error     string    `json:"error"`
	Causes    []string  `json:"causes"`
	Path      string    `json:"path"`
	Timestamp time.Time `json:"timestamp"`
}

func NewHttpException(message string, status int, causes []string) *HttpException {
	return &HttpException{
		Message: message,
		Status:  status,
		Causes:  causes,
	}
}
