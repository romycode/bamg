package noter

import "context"

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
		All(ctx context.Context) []UserInfo
		Save(ctx context.Context, u User)
		Delete(ctx context.Context, id string)
	}
)

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=UserRepository
