package sqlite

const (
	sqlUsersTable = "users"
)

type sqlUsers struct {
	ID     string `db:"id"`
	UserID string `db:"userId"`
	IBAN   string `db:"iban"`
	Credit string `db:"credit"`
}
