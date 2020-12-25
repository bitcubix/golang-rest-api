package db

import "errors"

var (
	ErrDBConnectionFailed = errors.New("connection to database failed")
	ErrUnknownDriver      = errors.New("database driver unknown")
)
