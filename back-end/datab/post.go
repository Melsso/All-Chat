package datab

import (
	"All-Chat/back-end/models"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetPosts(userID int) ([]models.Post, error) {
	query := `
        SELECT post_id, user_id, content, created_at
        FROM posts
        ORDER BY created_at DESC
    `
	rows, err := Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		var createdAt string
		err := rows.Scan(&post.PostID, &post.UserID, &post.Content, &createdAt)
		if err != nil {
			return nil, err
		}
		post.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
		if err != nil {
			return nil, err
		}
		user, err := GetUserById(post.UserID)
		if err != nil {
			return nil, err
		}
		post.PostOwner = strings.ToUpper(user.FirstName) + " " + strings.ToUpper(user.LastName)
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func GetComments(postID int) ([]models.Comment, error) {
	query := `
        SELECT comment_id, user_id, post_id, content, created_at
        FROM comments WHERE post_id = ?
    `
	rows, err := Db.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		var createdAt string
		err := rows.Scan(&comment.CommentID, &comment.UserID, &comment.PostId, &comment.Content, &createdAt)
		if err != nil {
			return nil, err
		}
		comment.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
		if err != nil {
			return nil, err
		}
		user, err := GetUserById(comment.UserID)
		if err != nil {
			return nil, err
		}
		comment.CommentOwner = user.FirstName + " " + user.LastName
		comments = append(comments, comment)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

func CreatePost(post *models.Post) error {
	query := `
        INSERT INTO posts (user_id, content, created_at)
        VALUES (?, ?, ?)
    `
	currentTime := time.Now()
	_, err := Db.Exec(query, post.UserID, post.Content, currentTime)
	if err != nil {
		return err
	}
	post.CreatedAt = currentTime
	return nil
}

func LikePost(userID, postID int) error {
	query := `
        INSERT INTO likes (user_id, post_id)
        VALUES (?, ?)
        ON DUPLICATE KEY UPDATE created_at = CURRENT_TIMESTAMP
    `
	_, err := Db.Exec(query, userID, postID)
	return err
}

func CommentPost(userID, postID int, content string) error {
	query := `
        INSERT INTO comments (user_id, post_id, content)
        VALUES (?, ?, ?)
    `
	_, err := Db.Exec(query, userID, postID, content)
	return err
}
