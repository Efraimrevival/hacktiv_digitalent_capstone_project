package errs

import "net/http"

type MessageErr interface {
	Message() string
	Status() int
}

type messageErrData struct {
	ErrMessage string
	ErrStatus  int
}

func (e *messageErrData) Message() string {
	return e.ErrMessage
}

func (e *messageErrData) Status() int {
	return e.ErrStatus
}

func NewError(status int) MessageErr {
	return &messageErrData{
		ErrMessage: http.StatusText(status),
		ErrStatus:  status,
	}
}

func NewCustomError(status int, message string) MessageErr {
	return &messageErrData{
		ErrMessage: message,
		ErrStatus:  status,
	}
}
