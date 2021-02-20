package blogs

import "time"

// Blog is
type Blog struct {
	ID        string     `json:"id" bson:"id"`
	Title     string     `json:"title" bson:"title"`
	Body      string     `json:"body" bson:"body"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" bson:"deleted_at"`
}

// TableName is
func (Blog) TableName() string {
	return "blogs"
}
