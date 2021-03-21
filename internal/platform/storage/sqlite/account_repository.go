package sqlite

import (
	"context"
	"database/sql"
	"log"

	noter "github.com/romycode/bank-manager/internal"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) noter.AccountRepository {
	return &AccountRepository{db: db}
}

func (ar AccountRepository) GetByUserId(ctx context.Context, userID string) []noter.Account {
	rows, err := ar.db.QueryContext(ctx, "SELECT * FROM accounts WHERE user_id = ?;", userID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var accounts []noter.Account
	for rows.Next() {
		a := *new(noter.Account)
		err = rows.Scan(&a.ID, &a.UserID, &a.IBAN, &a.Credit)
		if err != nil {
			log.Fatal(err)
		}
		accounts = append(accounts, a)
	}
	return accounts
}

func (ar AccountRepository) All(ctx context.Context) []noter.Account {
	rows, err := ar.db.QueryContext(ctx, "SELECT * FROM accounts;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var accounts []noter.Account
	for rows.Next() {
		a := noter.Account{}
		err := rows.Scan(&a.ID, &a.UserID, &a.IBAN, &a.Credit)
		if err != nil {
			log.Fatal(err)
		}
		accounts = append(accounts, a)
	}
	return accounts
}

func (ar AccountRepository) Save(ctx context.Context, a noter.Account) {
	stmt, err := ar.db.PrepareContext(ctx, "INSERT INTO accounts VALUES (?, ?, ?, ?);")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.ExecContext(ctx, a.ID, a.UserID, a.IBAN, a.Credit)
	if err != nil {
		log.Fatal(err)
	}
}

func (ar AccountRepository) Delete(ctx context.Context, id string) {
	stmt, err := ar.db.PrepareContext(ctx, "DELETE FROM accounts WHERE id = ?;")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
}
