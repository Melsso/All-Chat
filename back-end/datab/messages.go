package datab

import (
	"database/sql"
	"All-Chat/back-end/models"
	"time"
	"fmt"
	
	_ "github.com/go-sql-driver/mysql"
)

func StartOrGetConversation(userid, friendid int) (int, error) {
	var conversationID int
	query := `
        SELECT conversation_id 
        FROM conversations 
        WHERE (user1_id = ? AND user2_id = ?) OR (user1_id = ? AND user2_id = ?)
    `
	err := Db.QueryRow(query, userid, friendid, friendid, userid).Scan(&conversationID)
	if err == sql.ErrNoRows {
		iquery := `
		INSERT INTO conversations (user1_id, user2_id) 
            VALUES (?, ?)
		`
		result, err := Db.Exec(iquery, userid, friendid)
		if err != nil {
			return 0, err
		}
		newConvID, err := result.LastInsertId()
		if err != nil {
			return 0, err
		}
		return int(newConvID), nil
	} else if err != nil {
		return 0, err
	}
	return conversationID, nil
}

func GetMessages(conversationID int) ([]models.Message, error) {
	query := `
        SELECT message_id, sender_id, content, created_at, is_read
        FROM messages
        WHERE conversation_id = ?
        ORDER BY created_at
    `
	rows, err := Db.Query(query, conversationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var message models.Message
		var createdAtRaw []uint8

		err := rows.Scan(&message.MessageID, &message.SenderID, &message.Content, &createdAtRaw, &message.IsRead)
		if err != nil {
			return nil, err
		}

		createdAtString := string(createdAtRaw)
		message.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtString)
        if err != nil {
            return nil, fmt.Errorf("failed to parse created_at: %v", err)
        }
		messages = append(messages, message)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return messages, nil
}

func SendMessage(conversationID, senderID int, content string) error {
	query := `
        INSERT INTO messages (conversation_id, sender_id, content)
        VALUES (?, ?, ?)
    `
	_, err := Db.Exec(query, conversationID, senderID, content)
	if err != nil {
		return err
	}
	return nil
}
