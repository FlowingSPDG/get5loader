package database

import (
	"errors"
)

var (
	ErrNotFound = errors.New("specified resource not found")
	ErrInternal = errors.New("internal error")
)

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

func IsInternal(err error) bool {
	return errors.Is(err, ErrInternal)
}
