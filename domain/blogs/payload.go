package blogs

import "time"

// CreateBlog is
type CreateBlog struct {
	ID    string `json:"id" bson:"id"`
	Title string `json:"title" bson:"title"`
	Body  string `json:"body" bson:"body"`
}

// UpdateBlog is
type UpdateBlog struct {
	ID        string    `json:"id" bson:"id"`
	Title     string    `json:"title" bson:"title"`
	Body      string    `json:"body" bson:"body"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
