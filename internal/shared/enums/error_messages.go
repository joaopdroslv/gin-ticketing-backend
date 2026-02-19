package enums

type ErrorMessage string

const (
	InternalServerError ErrorMessage = "something went wrong internally"
	ResourceNotFound    ErrorMessage = "requested resource not found"
	ZeroRowsAffected    ErrorMessage = "nothing affected with the action"
	InvalidID           ErrorMessage = "the provided id is invalid"
	BadRequest          ErrorMessage = "bad request, fix it and try again later"
)
