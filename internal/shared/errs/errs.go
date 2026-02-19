package errs

import "errors"

var (
	// General
	ErrZeroRowsReturned = errors.New("zero rows returned")
	ErrResourceNotFound = errors.New("resource not found")
	ErrNothingToUpdate  = errors.New("nothing to update")
	ErrZeroRowsAffected = errors.New("zero rows affected")
	ErrValidationError  = errors.New("validation error")

	// Auth
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden")

	// Account
	ErrInactiveAccount = errors.New("account is inactive, cannot perform this action, request its reactivation")
	ErrDeletedAccount  = errors.New("deleted account, cannot perform this action")
)
