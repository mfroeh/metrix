package lolapi

import (
	"errors"
	"net/http"
)

// import (
// 	"errors"
// 	"fmt"
// )

// type UnknownError struct {
// 	Err error
// }

// func (we *UnknownError) Error() string {
// 	if we.Err != nil {
// 		return fmt.Sprintf("lolapi: %s", we.Err.Error())
// 	}
// 	return "lolapi: unknown error"
// }

// func (we *UnknownError) Unwrap() error {
// 	return we.Err
// }

// func NewUnknownError(err error) *UnknownError {
// 	return &UnknownError{
// 		Err: err,
// 	}
// }

var (
	ErrTooManyRequests   = errors.New("lolapi: too many requests")
	ErrInvalidRouteOrKey = errors.New("lolapi: invalid route or api key")
	ErrNotAuthorized     = errors.New("lolapi: not authorized")
	ErrResourceNotFound  = errors.New("lolapi: resource not found")
	ErrBadRequest        = errors.New("lolapi: bad request")
)

var ErrorForStatusCode = map[int]error{
	http.StatusTooManyRequests: ErrTooManyRequests,
	http.StatusUnauthorized:    ErrNotAuthorized,
	http.StatusForbidden:       ErrInvalidRouteOrKey,
	http.StatusNotFound:        ErrResourceNotFound,
	http.StatusBadRequest:      ErrBadRequest,
}
