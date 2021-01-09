package repositories

import (
	"fmt"

	"github.com/romycode/bank-manager/database"

	"github.com/romycode/bank-manager/models"
)

var db = database.GetConnection()

func GetAllUsers() []models.UserInfo {
	rows, err := db.Query("SELECT * FROM users;", nil)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var users []models.UserInfo
	for rows.Next() {
		u := *new(models.User)
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			fmt.Println(err)
		}
		accounts := GetAccountByUserId(u.ID)
		users = append(users, models.UserInfo{
			User:     u,
			Accounts: accounts,
		})
	}
	return users
}

func SaveUser(u *models.User) {
	stmt, err := db.Prepare("INSERT INTO users VALUES (?, ?, ?);")
	if err != nil {
		fmt.Println(err)
	}
	_, err = stmt.Exec(u.ID, u.Name, u.Email)
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteUser(id string) {
	stmt, _ := db.Prepare("DELETE FROM users WHERE id = ?;")
	_, err := stmt.Exec(id)
	if err != nil {
		fmt.Println(err)
	}
}
