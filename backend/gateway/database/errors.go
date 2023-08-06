package database

import (
	"errors"

	"golang.org/x/xerrors"
)

var (
	ErrNotFound = xerrors.New("specified resource not found")
	ErrInternal = xerrors.New("internal error")
)

func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

func IsInternal(err error) bool {
	return errors.Is(err, ErrInternal)
}
