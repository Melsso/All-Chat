package datab

import (
	"database/sql"
	"log"
	"All-Chat/back-end/models"

	_ "github.com/go-sql-driver/mysql"
)

func InsertUser(regForm models.RegistrationForm) (sql.Result, error) {
	query := `
        INSERT INTO users (first_name, last_name, user_name, date_of_birth, email, password, gender, user_name)
        VALUES (?, ?, ?, ?, ?, ?, ?, CONCAT(?, user_id))
    `
	result, err := Db.Exec(query, regForm.FirstName, regForm.LastName, regForm.DateOfBirth, regForm.Email, regForm.Password, regForm.Gender)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}
