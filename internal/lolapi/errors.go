package lolapi

import "errors"

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
	ErrTooManyRequests = errors.New("lolapi: too many requests")
)
