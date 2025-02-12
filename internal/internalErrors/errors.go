package internalerrors

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrInternal  = errors.New("internal server error")
	ErrNotFound  = errors.New("resource not found")
	ErrInvalidID = errors.New("invalid ID provided")
)

func ProcessErrorToReturn(err error) error {
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrInternal
	}
	return err
}
