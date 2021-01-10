package repositories

import (
	"database/sql"

	"github.com/romycode/bank-manager/errors"

	"github.com/romycode/bank-manager/models"
)

type UserRepository interface {
	All() []models.UserInfo
	Save(u *models.User)
	Delete(id string)
}

type SqliteUserRepository struct {
	db *sql.DB
	ac AccountRepository
}

func NewSqliteUserRepository(db *sql.DB) UserRepository {
	return SqliteUserRepository{db: db, ac: NewSqliteAccountRepository(db)}
}

func (ur SqliteUserRepository) All() []models.UserInfo {
	rows, err := ur.db.Query("SELECT * FROM users;", nil)
	errors.HandleError(err)
	defer rows.Close()

	var users []models.UserInfo
	for rows.Next() {
		u := *new(models.User)
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		errors.HandleError(err)
		accounts := ur.ac.GetByUserId(u.ID)
		users = append(users, models.UserInfo{
			User:     u,
			Accounts: accounts,
		})
	}
	return users
}

func (ur SqliteUserRepository) Save(u *models.User) {
	stmt, err := ur.db.Prepare("INSERT INTO users VALUES (?, ?, ?);")
	errors.HandleError(err)
	_, err = stmt.Exec(u.ID, u.Name, u.Email)
	errors.HandleError(err)
}

func (ur SqliteUserRepository) Delete(id string) {
	stmt, _ := ur.db.Prepare("DELETE FROM users WHERE id = ?;")
	_, err := stmt.Exec(id)
	errors.HandleError(err)
}
