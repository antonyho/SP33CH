package model

import "time"

// Post is the root of all discussions
// A post must have at least one topic
// User can reply upon post
type Post struct {
	Topic     string    `json:"topic"`   // multiple?
	Subject   string    `json:"subject"` // Length restriction should be applied
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
