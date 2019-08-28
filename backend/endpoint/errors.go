package endpoint

import (
	"errors"
	"net/http"
	"strings"
)

var (
	// ErrInternalError for internal error
	ErrInternalError = errors.New("internal error")

	// ErrInvalidBody for invalid request body
	ErrInvalidBody = errors.New("invalid request body")

	// ErrBadRequest for invalid request
	ErrBadRequest = errors.New("bad request")

	// ErrNotFound for not found
	ErrNotFound = errors.New("not found")
)

// HandleErrors ...
func HandleErrors(err error) (int, error) {
	if strings.Contains(err.Error(), "not found") {
		return http.StatusNotFound, ErrNotFound
	}
	if strings.Contains(err.Error(), "bad request") || strings.Contains(err.Error(), "invalid request") {
		return http.StatusBadRequest, ErrBadRequest
	}
	if strings.Contains(err.Error(), "invalid body") {
		return http.StatusBadRequest, ErrInvalidBody
	}

	return http.StatusInternalServerError, ErrInternalError
}
