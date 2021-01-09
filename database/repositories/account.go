package repositories

import (
	"fmt"

	"github.com/romycode/bank-manager/database"
	"github.com/romycode/bank-manager/models"
)

var db = database.GetConnection()

func GetAllAccounts() []models.Account {
	rows, err := db.Query("SELECT * FROM accounts;", nil)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var accounts []models.Account
	for rows.Next() {
		a := *new(models.Account)
		err := rows.Scan(&a.ID, &a.UserID, &a.IBAN, &a.Credit)
		if err != nil {
			fmt.Println(err)
		}
		accounts = append(accounts, a)
	}
	return accounts
}

func GetAccountByUserId(userId string) []models.Account {
	rows, err := db.Query("SELECT * FROM accounts WHERE user_id = ?;", userId)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var accounts []models.Account
	for rows.Next() {
		a := *new(models.Account)
		err = rows.Scan(&a.ID, &a.UserID, &a.IBAN, &a.Credit)
		if err != nil {
			fmt.Println(err)
		}
		accounts = append(accounts, a)
	}
	return accounts
}

func SaveAccount(a *models.Account) {
	stmt, err := db.Prepare("INSERT INTO accounts VALUES (?, ?, ?, ?);")
	if err != nil {
		fmt.Println(err)
	}
	_, err = stmt.Exec(a.ID, a.UserID, a.IBAN, a.Credit)
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteAccount(id string) {
	stmt, err := db.Prepare("DELETE FROM accounts WHERE id = ?;")
	if err != nil {
		fmt.Println(err)
	}
	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
}
