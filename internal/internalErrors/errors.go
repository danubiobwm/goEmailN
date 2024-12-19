package internalerrors

import "errors"

var (
	ErrInternal  = errors.New("internal server error")
	ErrNotFound  = errors.New("resource not found")
	ErrInvalidID = errors.New("invalid ID provided")
)
