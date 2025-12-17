package errors

import "errors"

var (
	ErrNotFound        = errors.New("resource not found")
	ErrNothingToUpdate = errors.New("nothing to update")
	ErrConflict        = errors.New("resource conflict")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrForbidden       = errors.New("forbidden")
	ErrValidation      = errors.New("validation error")
)
