package jsonplaceholder

// Post is
type Post struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
