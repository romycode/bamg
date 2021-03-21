package sqlite

const (
	sqlAccountTable = "notes"
)

type sqlAccount struct {
	ID     string `db:"id"`
	UserID string `db:"userId"`
	IBAN   string `db:"iban"`
	Credit string `db:"credit"`
}
