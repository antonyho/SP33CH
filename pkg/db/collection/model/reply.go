package model

import "time"

// Reply on a post as a discussion tree
type Reply struct {
	PostID    string    `json:"post_id"`
	ReplyID   string    `json:"reply_id"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
