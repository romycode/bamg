package sqlite

import (
	"context"
	"database/sql"
	"testing"

	noter "github.com/romycode/bank-manager/internal"

	"github.com/stretchr/testify/assert"

	"github.com/romycode/bank-manager/internal/bank_manager_api/database"
)

func TestSqliteUserRepository_Delete(t *testing.T) {
	ur := UserRepository{
		db: database.GetConnection(),
	}

	userID := "aass212"

	_, _ = ur.db.Exec(`
		DELETE FROM users;
		DELETE FROM accounts;

		INSERT INTO users(id, name, email) 
			VALUES(?, 'test', 'test@test.com');
	`, userID)

	ur.Delete(context.Background(), userID)

	rows, _ := ur.db.Query("SELECT * FROM users WHERE id = ?", userID)

	var count int
	for rows.Next() {
		count++
	}

	assert.Equal(t, count, 0)
}

func TestSqliteUserRepository_Save(t *testing.T) {
	ur := UserRepository{
		db: database.GetConnection(),
	}

	userID := "1234"

	_, _ = ur.db.Exec(`
		DELETE FROM users;
		DELETE FROM accounts;

		INSERT INTO users(id, name, email) 
			VALUES(?, 'test', 'test@test.com');
	`, userID)

	u := noter.User{
		ID:    userID,
		Name:  "t3stname",
		Email: "ES00000000000000",
	}

	ur.Save(context.Background(), u)

	rows, _ := ur.db.Query("SELECT * FROM users WHERE id = ?", userID)
	var count int
	for rows.Next() {
		count++
	}

	assert.Equal(t, count, 1)
}

func TestSqliteUserRepository_All(t *testing.T) {
	type fields struct {
		db *sql.DB
		ac noter.AccountRepository
	}
	tests := []struct {
		name      string
		insertSQL string
		fields    fields
		want      []noter.UserInfo
	}{
		{
			name: "Return 0",
			insertSQL: `
				DELETE FROM accounts;
				DELETE FROM users;
			`,
			fields: fields{
				db: database.GetConnection(),
				ac: NewAccountRepository(database.GetConnection()),
			},
			want: nil,
		},
		{
			name: "Return 1",
			insertSQL: `
				DELETE FROM accounts;
				DELETE FROM users;

				INSERT INTO users(id, name, email) 
					VALUES('sddkl-324-ldf', 'test', 'test@test.com');
			`,
			fields: fields{
				db: database.GetConnection(),
				ac: NewAccountRepository(database.GetConnection()),
			},
			want: []noter.UserInfo{
				{
					User: noter.User{
						ID:    "sddkl-324-ldf",
						Name:  "test",
						Email: "test@test.com",
					},
					Accounts: nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _ = tt.fields.db.Exec(tt.insertSQL)

			ur := NewUserRepository(tt.fields.db, tt.fields.ac)

			res := ur.All(context.Background())

			assert.EqualValues(t, res, tt.want)
		})
	}
}
