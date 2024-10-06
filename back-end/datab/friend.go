package datab

import (
	"database/sql"
	"All-Chat/back-end/models"

	_ "github.com/go-sql-driver/mysql"
)

func Deletefriend(userid int, friendid int) error {
	query := `
        DELETE FROM friends 
        WHERE (user_id = ? AND friend_id = ?) OR (user_id = ? AND friend_id = ?)
	`
	_, err := Db.Exec(query, userid, friendid, friendid, userid)
	return err
}

func Acceptfriend(userid int, friendid int) error {
	query := `
		UPDATE friends 
		SET status = ?
		WHERE user_id = ? AND friend_id = ? AND status = ?
	`
	_, err := Db.Exec(query, models.Accepted, friendid, userid, models.Pending)
	return err
}

func RemoveFriendReq(userid int, friendid int) error {
	query := `
		DELETE FROM friends
		WHERE friend_id = ? AND user_id = ? AND status = ?
	`
	_, err := Db.Exec(query, friendid, userid, models.Pending)
	return err
}

func Addfriend(userid int, friendid int) (string, error) {
	var status string
	query := `SELECT status 
		FROM friends 
		WHERE (user_id = ? AND friend_id = ?) 
		OR (friend_id = ? AND user_id = ?)
	`
	err := Db.QueryRow(query, userid, friendid, friendid, userid).Scan(&status)
	if err == sql.ErrNoRows {

		iquery := `
			INSERT INTO friends (user_id, friend_id, status) 
			VALUES (?, ?, ?)
		`
		_, err = Db.Exec(iquery, userid, friendid, models.Pending)
		if err != nil {
			return "", err
		}
		return "pending", err
	}
	
	if err != nil {
		return "", err
	}

	if status == "pending" || status == "accepted" {
		return status, nil
	}


	return "", nil
}

func GetInvites(userid int) ([]models.User, error) {
	query := `
		SELECT u.user_id, u.first_name, u.last_name, u.date_of_birth, u.email, u.gender
		FROM users u JOIN friends f ON u.user_id = f.friend_id
		WHERE f.user_id = ? AND f.status = 'pending'
	`
	rows, err := Db.Query(query, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invite_list []models.User

	for rows.Next() {
		var person models.User
		err := rows.Scan(&person.UserID, &person.FirstName, &person.LastName, &person.DateOfBirth, &person.Email, &person.Gender)
		if err != nil {
			return nil, err
		}
		invite_list = append(invite_list, person)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return invite_list, nil
}

func GetFriends(userid int) ([]models.User, error) {
	query := `
		SELECT u.user_id, u.first_name, u.last_name, u.date_of_birth, u.email, u.gender
		FROM users u
		JOIN friends f ON (u.user_id = f.friend_id OR u.user_id = f.user_id)
		WHERE (f.user_id = ? OR f.friend_id = ?) AND f.status = 'accepted'
		AND u.user_id != ?
	`
	rows, err := Db.Query(query, userid, userid, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var friends []models.User

	for rows.Next() {
		var friend models.User
		err := rows.Scan(&friend.UserID, &friend.FirstName, &friend.LastName, &friend.DateOfBirth, &friend.Email, &friend.Gender)
		if err != nil {
			return nil, err
		}
		friends = append(friends, friend)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return friends, nil
}
