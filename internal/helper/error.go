package helper

import (
	"errors"
	"net/http"
)

type Error struct {
	Message  string
	Error    string
	Code     string
	HttpCode int
}

func (e Error) ErrorMessage() string {
	return e.Message
}

func NewError(msg string, err string, code string, httpCode int) Error {
	return Error{
		Message:  msg,
		Error:    err,
		Code:     code,
		HttpCode: httpCode,
	}
}

var (
	ErrNotFound   = errors.New("not found")
	ErrBadRequest = errors.New("bad request")
)

var (
	ErrTitleRequired = errors.New("title is required")
	ErrLinkRequired  = errors.New("link is required")
	ErrNoteNotFound  = errors.New("note is not found")
)

var (
	ErrorTitleRequired = NewError(ErrBadRequest.Error(), ErrTitleRequired.Error(), "40001", http.StatusBadRequest)
	ErrorLinkRequired  = NewError(ErrBadRequest.Error(), ErrLinkRequired.Error(), "40002", http.StatusBadRequest)
	ErrorNoteNotFound  = NewError(ErrNotFound.Error(), ErrNotFound.Error(), "40401", http.StatusNotFound)
	ErrorGeneral       = NewError("internal server error", "unknown error", "99999", http.StatusInternalServerError)
)

var (
	ErrorMapping = map[string]Error{
		ErrTitleRequired.Error(): ErrorTitleRequired,
		ErrLinkRequired.Error():  ErrorLinkRequired,
		ErrNoteNotFound.Error():  ErrorNoteNotFound,
	}
)
