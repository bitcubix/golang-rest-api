package db

import "errors"

var (
	ErrUnknownDriver = errors.New("database driver unknown")
)
