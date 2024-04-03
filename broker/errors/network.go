package errors

import "errors"

var ErrValidationFailed = errors.New("validation failed")

var ErrTimeout = errors.New("timeout")
