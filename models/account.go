package models

import (
	"fmt"
	"time"
)

type Account struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	IBAN   string `json:"iban"`
	Credit string `json:"credit"`
}

// NewIban create a new fake IBAN number
func NewIban() string {
	t := time.Now().UnixNano()
	return fmt.Sprintf("ES%02d%04d%016d", t%100, t%10000, t%10000000000000000)
}
