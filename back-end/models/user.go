package models

import "time"

type User struct {
	UserID      int    `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	Email       string `json:"email"`
	Password    string `json:"-"`
	Gender      string `json:"gender"`
}

type RegistrationForm struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"` // Should be parsed into time.Time
	Email       string `json:"email"`
	Password    string `json:"password"`
	Gender      string `json:"gender"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Post struct {
	PostOwner string	`json:"post_owner"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type Comment struct {
	CommentID		int 		`json:"comment_id"`
	UserID			int			`json:"user_id"`
	CommentOwner 	string		`json:"comment_owner"`
	PostId			int			`json:"post_id"`
	Content			string		`json:"content"`	
	CreatedAt 		time.Time	`json:"created_at"`
}

type FriendStatus string

const (
	Pending		FriendStatus = "pending"
	Accepted	FriendStatus = "accepted"
	Blocked		FriendStatus = "blocked"
)

type Friend struct {
	FriendshipID	int				`json:"friendship_id"`
	UserID			int				`json:"user_id"`
	FriendID		int				`json:"friend_id"`
	Status			FriendStatus	`json:"status" db:"status"`
	CreatedAt		time.Time		`json:"created_at"`
}

type Message struct {
    MessageID     int       `json:"message_id"`
    ConversationID int      `json:"conversation_id"`
    SenderID      int       `json:"sender_id"`
    Content       string    `json:"content"`
    CreatedAt     time.Time `json:"created_at"`
    IsRead        bool      `json:"is_read"`
}

type Conversation struct {
    ConversationID int       `json:"conversation_id"`
    User1ID        int       `json:"user1_id"`
    User2ID        int       `json:"user2_id"`
    CreatedAt      time.Time `json:"created_at"`
}

type MsgConversation struct {
	ConversationID	int `json:"conversation_id"`
	Content			string `json:"content"`
}