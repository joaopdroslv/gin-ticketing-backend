package enums

type Message string

const (
	ErrInternal   Message = "Something went wrong internally."
	ErrNotFound   Message = "Requested resource not found."
	ErrInvalidID  Message = "The provided ID is invalid"
	ErrBadRequest Message = "Bad request, fix it and try again later."
)
