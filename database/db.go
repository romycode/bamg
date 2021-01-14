package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3" // Driver to connect to SQLite3

	"github.com/romycode/bank-manager/errors"
)

var db *sql.DB

func GetConnection() *sql.DB {
	if db != nil {
		return db
	}
	path := os.Getenv("BANK_MANAGER")
	sqlite, err := sql.Open("sqlite3", path+"/data.db")
	errors.HandleError(err)

	createDatabase(sqlite)
	return sqlite
}

func createDatabase(sqlite *sql.DB) {
	var err error

	usersTable := `
	CREATE TABLE IF NOT EXISTS users( 
	    id VARCHAR(36) PRIMARY KEY, 
	    name VARCHAR(50), 
	    email VARCHAR(50)
	);`

	accountsTable := `
	CREATE TABLE IF NOT EXISTS accounts( 
    	id VARCHAR(36) PRIMARY KEY, 
    	user_id VARCHAR(50), 
    	iban VARCHAR(50), 
    	credit DOUBLE(1000, 2), 
    	CONSTRAINT fk_user_id 
    	    FOREIGN KEY (user_id) REFERENCES users (user_id)
	);`

	_, err = sqlite.Exec(usersTable)
	errors.HandleError(err)

	_, err = sqlite.Exec(accountsTable)
	errors.HandleError(err)
}
