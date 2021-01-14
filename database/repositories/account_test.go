package repositories

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/romycode/bank-manager/database"
	"github.com/romycode/bank-manager/models"
)

func TestSqliteAccountRepository_All(t *testing.T) {
	ar := SqliteAccountRepository{
		db: database.GetConnection(),
	}

	_, _ = ar.db.Exec(`
		DELETE FROM accounts;
		DELETE FROM users;

		INSERT INTO users(id, name, email) 
			VALUES('sddkl-324-ldf', 'test', 'test@test.com');

		INSERT INTO accounts(id, user_id, iban, credit) 
			VALUES('1234', 'sddkl-324-ldf', 'ES00000000000000', '0');
	`)

	actual := ar.All()

	assert.Len(t, actual, 1)
}

func TestSqliteAccountRepository_Delete(t *testing.T) {
	ar := SqliteAccountRepository{
		db: database.GetConnection(),
	}

	accountID := "1234"

	_, _ = ar.db.Exec(`
		DELETE FROM accounts;
		DELETE FROM users;

		INSERT INTO users(id, name, email) 
			VALUES('aass212', 'test', 'test@test.com');

		INSERT INTO accounts(id, user_id, iban, credit) 
			VALUES(?, 'aass212', 'ES00000000000000', '0');
	`, accountID)

	ar.Delete(accountID)

	rows, _ := ar.db.Query("SELECT * FROM accounts WHERE id = ?", accountID)

	var count int
	for rows.Next() {
		count++
	}

	assert.Equal(t, count, 0)
}

func TestSqliteAccountRepository_GetByUserId(t *testing.T) {
	ar := SqliteAccountRepository{
		db: database.GetConnection(),
	}

	userID := "1234"

	_, _ = ar.db.Exec(`
		DELETE FROM accounts;
		DELETE FROM users;

		INSERT INTO users(id, name, email) 
			VALUES(?, 'test', 'test@test.com');

		INSERT INTO accounts(id, user_id, iban, credit) 
			VALUES('1234', ?, 'ES00000000000000', '0');
	`, userID, userID)

	actual := ar.GetByUserId(userID)

	assert.Len(t, actual, 1)
}

func TestSqliteAccountRepository_Save(t *testing.T) {
	ar := SqliteAccountRepository{
		db: database.GetConnection(),
	}

	_, _ = ar.db.Exec(`
		DELETE FROM accounts;
		DELETE FROM users;

		INSERT INTO users(id, name, email) 
			VALUES('sddkl-324-ldf', 'test', 'test@test.com');
	`)

	accountID := "1234"
	a := models.Account{
		ID:     accountID,
		UserID: "sddkl-324-ldf",
		IBAN:   "ES00000000000000",
		Credit: "0",
	}

	ar.Save(a)

	rows, _ := ar.db.Query("SELECT * FROM accounts WHERE id = ?", accountID)
	var count int
	for rows.Next() {
		count++
	}

	assert.Equal(t, count, 1)
}
