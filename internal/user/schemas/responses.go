package schemas

import sharedschemas "go-gin-ticketing-backend/internal/shared/schemas"

type ResponseUser struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Birthdate  string `json:"birthdate"`
	Email      string `json:"email"`
	UserStatus string `json:"user_status"`
}

type GetAllUsersResponse struct {
	Items      []ResponseUser                   `json:"items"`
	Pagination sharedschemas.ResponsePagination `json:"pagination"`
}

type DeleteUserResponse struct {
	ID      int64 `json:"id"`
	Deleted bool  `json:"deleted"`
}
