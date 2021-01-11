package repositories

import (
	"database/sql"

	"github.com/romycode/bank-manager/errors"
	"github.com/romycode/bank-manager/models"
)

type SqliteAccountRepository struct {
	db *sql.DB
}

func NewSqliteAccountRepository(db *sql.DB) models.AccountRepository {
	return SqliteAccountRepository{db: db}
}

func (ar SqliteAccountRepository) GetByUserId(userID string) []models.Account {
	rows, err := ar.db.Query("SELECT * FROM accounts WHERE user_id = ?;", userID)
	errors.HandleError(err)
	defer rows.Close()

	var accounts []models.Account
	for rows.Next() {
		a := *new(models.Account)
		err = rows.Scan(&a.ID, &a.UserID, &a.IBAN, &a.Credit)
		errors.HandleError(err)
		accounts = append(accounts, a)
	}
	return accounts
}

func (ar SqliteAccountRepository) All() []models.Account {
	rows, err := ar.db.Query("SELECT * FROM accounts;", nil)
	errors.HandleError(err)
	defer rows.Close()

	var accounts []models.Account
	for rows.Next() {
		a := *new(models.Account)
		err := rows.Scan(&a.ID, &a.UserID, &a.IBAN, &a.Credit)
		errors.HandleError(err)
		accounts = append(accounts, a)
	}
	return accounts
}

func (ar SqliteAccountRepository) Save(a models.Account) {
	stmt, err := ar.db.Prepare("INSERT INTO accounts VALUES (?, ?, ?, ?);")
	errors.HandleError(err)

	_, err = stmt.Exec(a.ID, a.UserID, a.IBAN, a.Credit)
	errors.HandleError(err)
}

func (ar SqliteAccountRepository) Delete(id string) {
	stmt, err := ar.db.Prepare("DELETE FROM accounts WHERE id = ?;")
	errors.HandleError(err)

	_, err = stmt.Exec(id)
	errors.HandleError(err)
}
