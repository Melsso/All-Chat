package datab

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func createUserTable() error {
	query := `CREATE TABLE IF NOT EXISTS users (
        user_id INT AUTO_INCREMENT PRIMARY KEY,
        first_name VARCHAR(50),
        last_name VARCHAR(50),
        date_of_birth DATE,
        email VARCHAR(100) UNIQUE,
        password VARCHAR(100),
        gender VARCHAR(10),
		username VARCHAR(100) -- Add the username field
    )`
	_, err := Db.Exec(query)
	return err
}

func createPostsTable() error {
	query := `CREATE TABLE IF NOT EXISTS posts (
        post_id INT AUTO_INCREMENT PRIMARY KEY,
        user_id INT,
        content TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (user_id) REFERENCES users(user_id)
    )`
	_, err := Db.Exec(query)
	return err
}

func createFriendsTable() error {
	query := `CREATE TABLE IF NOT EXISTS friends (
    	friendship_id INT AUTO_INCREMENT PRIMARY KEY,
    	user_id INT NOT NULL,
    	friend_id INT NOT NULL,
    	status ENUM('pending', 'accepted', 'blocked') NOT NULL,
    	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    	FOREIGN KEY (friend_id) REFERENCES users(user_id) ON DELETE CASCADE,
    	UNIQUE KEY (user_id, friend_id),
    	INDEX (user_id),
    	INDEX (friend_id)
	);`
	_, err := Db.Exec(query)
	return err
}

func createLikesTable() error {
	query := `CREATE TABLE IF NOT EXISTS likes (
        like_id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		post_id INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(user_id),
		FOREIGN KEY (post_id) REFERENCES posts(post_id)
    )`
	_, err := Db.Exec(query)
	return err
}

func createCommentsTable() error {
	query := `CREATE TABLE IF NOT EXISTS comments (
        comment_id INT AUTO_INCREMENT PRIMARY KEY,
		user_id INT NOT NULL,
		post_id INT NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(user_id),
		FOREIGN KEY (post_id) REFERENCES posts(post_id)
    )`
	_, err := Db.Exec(query)
	return err
}

func createConversationsTable() error {
	query := `CREATE TABLE IF NOT EXISTS conversations (
		conversation_id INT AUTO_INCREMENT PRIMARY KEY,
		user1_id INT NOT NULL,
		user2_id INT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UNIQUE KEY (user1_id, user2_id),
		CONSTRAINT fk_user1 FOREIGN KEY (user1_id) REFERENCES users(user_id) ON DELETE CASCADE,
		CONSTRAINT fk_user2 FOREIGN KEY (user2_id) REFERENCES users(user_id) ON DELETE CASCADE
	);`
	_, err := Db.Exec(query)
	return err
}

func createMessagesTable() error {
	query := `CREATE TABLE IF NOT EXISTS messages(
    	message_id INT AUTO_INCREMENT PRIMARY KEY,
    	conversation_id INT NOT NULL,
    	sender_id INT NOT NULL,
    	content TEXT NOT NULL,
    	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    	is_read BOOLEAN DEFAULT FALSE,
    	CONSTRAINT fk_conversation FOREIGN KEY (conversation_id) REFERENCES conversations(conversation_id) ON DELETE CASCADE,
    	CONSTRAINT fk_sender FOREIGN KEY (sender_id) REFERENCES users(user_id) ON DELETE CASCADE
	);`
	_, err := Db.Exec(query)
	return err
}

func InitDB() {
	var err error

	dsn := "serv:pswd@tcp(localhost:3306)/chatdb"
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error; ", err)
		log.Fatal(err)
	}
	_, err = Db.Exec(`CREATE DATABASE IF NOT EXISTS chatdb`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = Db.Exec(`USE chatdb`)
	if err != nil {
		log.Fatal(err)
	}
	errUT := createUserTable()
	errPT := createPostsTable()
	errFT := createFriendsTable()
	errLT := createLikesTable()
	errCT := createCommentsTable()
	errCVT := createConversationsTable()
	errMT := createMessagesTable()
	if errUT != nil || errPT != nil || errFT != nil || errLT != nil || errCT != nil || errCVT != nil || errMT != nil {
		log.Fatal("Failed to create a necessary database...")
	}
}

func CloseDB() {
	Db.Close()
}