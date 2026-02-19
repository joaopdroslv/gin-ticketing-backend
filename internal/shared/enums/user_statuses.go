package enums

type UserStatusID int64

const (
	Active            UserStatusID = 1
	Inactive          UserStatusID = 2
	EmailConfirmation UserStatusID = 3
	DeletedAccount    UserStatusID = 4
)
