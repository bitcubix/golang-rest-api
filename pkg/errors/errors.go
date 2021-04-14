package errors

import "github.com/pkg/errors"

func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}
