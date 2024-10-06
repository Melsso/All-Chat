package datab

import (
	"database/sql"
	"All-Chat/back-end/models"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)

func CheckEmailExists(email string) (bool, error) {
	var userid int

	query := "SELECT user_id FROM users WHERE email = ?"
	err := Db.QueryRow(query, email).Scan(&userid)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := `
		SELECT user_id, first_name, last_name, date_of_birth, email, password, gender
		FROM users WHERE email = ?
	`
	err := Db.QueryRow(query, email).Scan(&user.UserID,
		&user.FirstName,
		&user.LastName,
		&user.DateOfBirth,
		&user.Email,
		&user.Password,
		&user.Gender)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserById(userID int) (models.User, error) {
	var user models.User
	query := `
		SELECT user_id, first_name, last_name, date_of_birth, email, password, gender
		FROM users WHERE user_id = ?
	`
	err := Db.QueryRow(query, userID).Scan(&user.UserID,
		&user.FirstName,
		&user.LastName,
		&user.DateOfBirth,
		&user.Email,
		&user.Password,
		&user.Gender)
	if err != nil {
		return user, err
	}
	return user, nil
}

func LookupUser(username string) ([]models.User, error) {
	var users []models.User

	words := strings.Fields(username)
	
	if len(words) == 0 {
		return nil, nil
	}

	var query string
	var args []interface{}

	if len(words) == 1 {
		query = `
			SELECT user_id, first_name, last_name
			FROM users
			WHERE first_name LIKE ? OR last_name LIKE ?
		`
		args = []interface{}{"%" + words[0] + "%", "%" + words[0] + "%"}
	} else {
		lname := words[0]
		fname := strings.Join(words[1:], " ")
		query = `
			SLECT user_id, first_name, last_name
			FROM users
			WHERE (first_name LIKE ? AND last_name LIKE ?)
			OR (first_name LIKE ? AND last_name LIKE ?)
		`
		args = []interface{}{
			"%" + fname + "%", "%" + lname + "%",
			"%" + lname + "%", "%" + fname + "%",
		}
	}

	rows, err := Db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.UserID, &user.FirstName, &user.LastName); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
