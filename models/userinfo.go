package models

type UserInfo struct {
	User
	Accounts []Account `json:"accounts"`
}
