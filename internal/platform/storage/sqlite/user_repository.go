package sqlite

import (
	"context"
	"database/sql"

	"log"

	noter "github.com/romycode/bank-manager/internal"
)

type UserRepository struct {
	db *sql.DB
	ac noter.AccountRepository
}

func NewUserRepository(db *sql.DB, ac noter.AccountRepository) noter.UserRepository {
	return &UserRepository{db: db, ac: ac}
}

func (ur *UserRepository) All(ctx context.Context) []noter.UserInfo {
	rows, err := ur.db.QueryContext(ctx, "SELECT * FROM users;", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []noter.UserInfo
	for rows.Next() {
		u := *new(noter.User)
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			log.Fatal(err)
		}
		accounts := ur.ac.GetByUserId(ctx, u.ID)
		if len(accounts) == 0 {
			accounts = make([]noter.Account, 1)
		}
		users = append(users, noter.UserInfo{
			User:     u,
			Accounts: accounts,
		})
	}
	return users
}

func (ur *UserRepository) Save(ctx context.Context, u noter.User) {
	stmt, err := ur.db.PrepareContext(ctx, "INSERT INTO users VALUES (?, ?, ?);")
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.ExecContext(ctx, u.ID, u.Name, u.Email)
	if err != nil {
		log.Fatal(err)
	}
}

func (ur *UserRepository) Delete(ctx context.Context, id string) {
	stmt, _ := ur.db.PrepareContext(ctx, "DELETE FROM users WHERE id = ?;")
	_, err := stmt.ExecContext(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
}
