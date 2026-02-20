package schemas

type ResponseUser struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Birthdate  string `json:"birthdate"`
	UserStatus string `json:"user_status"`
}

type GetAllUsersResponse struct {
	Total int64          `json:"total"`
	Users []ResponseUser `json:"users"`
}

type DeleteUserResponse struct {
	ID      int64 `json:"id"`
	Deleted bool  `json:"deleted"`
}
