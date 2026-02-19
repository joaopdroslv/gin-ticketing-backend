package errs

import "errors"

var (
	ErrResourceNotFound = errors.New("resource not found")
	ErrNothingToUpdate  = errors.New("nothing to update")
	ErrZeroRowsAffected = errors.New("zero rows affected")
	ErrResourceConflict = errors.New("resource conflict")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrForbidden        = errors.New("forbidden")
	ErrValidationError  = errors.New("validation error")
)
