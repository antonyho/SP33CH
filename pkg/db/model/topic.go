package model

// Topic must be given during post creation
// This classifies discussion type and attribute
type Topic struct {
	Name string `json:"name"` // Length restriction should be applied
}
