package models

type (
	User struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	UserInfo struct {
		User
		Accounts []Account `json:"accounts"`
	}

	UserRepository interface {
		All() []UserInfo
		Save(u *User)
		Delete(id string)
	}
)
