package domain

import "errors"

var (
	// Queries
	ErrZeroRowsReturned = errors.New("zero rows returned")
	ErrNothingToUpdate  = errors.New("nothing to update")
	ErrZeroRowsAffected = errors.New("zero rows affected")

	ErrResourceNotFound      = errors.New("resource not found")
	ErrResourceAlreadyExists = errors.New("this resource already exists")

	// User
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")

	// Auth
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrForbidden          = errors.New("forbidden")

	// User Status
	ErrInactiveUser                 = errors.New("inactive user, cannot login")
	ErrUserEmailConfirmationPending = errors.New("email confirmation is pending, do it before loggin in")
	ErrUserPasswordCreationPending  = errors.New("password creation is pending, do it before loggin in")
	ErrDeletedUser                  = errors.New("deleted user, cannot login")
)

// TODO: This could be improved to avoid multiple errors.Is conditions
func IsUserStatusRelated(err error) bool {
	return errors.Is(err, ErrInactiveUser) ||
		errors.Is(err, ErrUserEmailConfirmationPending) ||
		errors.Is(err, ErrUserPasswordCreationPending) ||
		errors.Is(err, ErrDeletedUser)
}
