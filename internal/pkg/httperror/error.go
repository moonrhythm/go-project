package httperror

import (
	"encoding/json"
	"net/http"
)

// Error is the http error
type Error struct {
	statusCode int
	Message    string
}

// New creates new http error
func New(statusCode int, message string) *Error {
	return &Error{
		statusCode: statusCode,
		Message:    message,
	}
}

// Error returns error message
func (err *Error) Error() string {
	return err.Message
}

// StatusCode returns http status code
func (err *Error) StatusCode() int {
	if err.statusCode == 0 {
		return http.StatusInternalServerError
	}
	return err.statusCode
}

func (err *Error) internal() bool {
	return err.StatusCode() == http.StatusInternalServerError
}

// MarshalJSON marshals error into json
func (err *Error) MarshalJSON() ([]byte, error) {
	var x struct {
		Error struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	if err.internal() {
		x.Error.Message = "Internal Server Error"
	} else {
		x.Error.Message = err.Message
	}
	return json.Marshal(x)
}
