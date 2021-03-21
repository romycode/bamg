package noter

import (
	"context"
	"fmt"
	"time"
)

type (
	Account struct {
		ID     string `json:"id"`
		UserID string `json:"user_id"`
		IBAN   string `json:"iban"`
		Credit string `json:"credit"`
	}

	AccountRepository interface {
		All(ctx context.Context) []Account
		GetByUserId(ctx context.Context, usrID string) []Account
		Save(ctx context.Context, a Account)
		Delete(ctx context.Context, id string)
	}
)

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=AccountRepository

// NewIban create a new fake IBAN number
func NewIban() string {
	t := time.Now().UnixNano()
	return fmt.Sprintf("ES%02d%04d%016d", t%100, t%10000, t%10000000000000000)
}
