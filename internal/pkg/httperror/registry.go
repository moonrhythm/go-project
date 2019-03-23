package httperror

import (
	"net/http"
	"sync"
)

var (
	mu     sync.Mutex
	errors = make(map[error]*Error)
)

// Register registers error with http error
// Register must be called while bootstrapping application or might panic with race
func Register(err error, httpErr *Error) {
	mu.Lock()
	errors[err] = httpErr
	mu.Unlock()
}

// From gets http error from error
func From(err error) *Error {
	if err, ok := err.(*Error); ok {
		return err
	}

	if err, ok := errors[err]; ok {
		return err
	}

	return New(http.StatusInternalServerError, err.Error())
}
