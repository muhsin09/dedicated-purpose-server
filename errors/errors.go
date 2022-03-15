package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Code    int
	Message string
}

var (
	// ErrInternal HTTP 500
	ErrInternal = &Error{
		Code:    http.StatusInternalServerError,
		Message: "Something went wrong",
	}
	// ErrBadRequest HTTP 400
	ErrBadRequest = &Error{
		Code:    http.StatusBadRequest,
		Message: "Error invalid argument",
	}
	// ErrEventNotFound HTTP 404
	ErrEventNotFound = &Error{
		Code:    http.StatusNotFound,
		Message: "Event not found",
	}
	// ErrMessageIsRequired HTTP 400
	ErrMessageIsRequired = &Error{
		Code:    http.StatusBadRequest,
		Message: "The message should be provided",
	}
	// ErrMessageSize HTTP 400
	ErrMessageSize = &Error{
		Code:    http.StatusBadRequest,
		Message: "The message size too much",
	}
)

func (err *Error) Error() string {
	return err.String()
}

func (err *Error) String() string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("error: code=%s message=%s", http.StatusText(err.Code), err.Message)
}

// JSON convert Error in json
func (err *Error) JSON() []byte {
	if err == nil {
		return []byte("{}")
	}
	res, _ := json.Marshal(err)
	return res
}

// StatusCode get status code
func (err *Error) StatusCode() int {
	if err == nil {
		return http.StatusOK
	}
	return err.Code
}
