package datab

import (
	"database/sql"
	"playground/models"

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

	query := `
		SELECT user_id, first_name, last_name ,user_name
		FROM users
		WHERE user_name LIKE ?
	`
	rows, err := Db.Query(query, "%"+username+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user models.User
		var usrn string
		if err := rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &usrn); err != nil {
			return nil, err
        }
		users = append(users, user)
	}
	return users, nil
}