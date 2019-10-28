package errors

import "github.com/pkg/errors"

func Errorf(format string, args... interface{}) error {
	return errors.Errorf(format, args...)
}
