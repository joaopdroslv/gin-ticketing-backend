package schemas

type CreateUserBody struct {
	Name      string `json:"name" binding:"required"`
	Birthdate string `json:"birthdate" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

type UpdateUserBody struct {
	Name      *string `json:"name" binding:"omitempty,min=2"`
	Email     *string `json:"email" binding:"omitempty,email"`
	Birthdate *string `json:"birthdate" binding:"omitempty,email"`
}
