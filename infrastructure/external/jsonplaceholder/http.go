package jsonplaceholder

// JPH is
type JPH struct {
	Posts *Posts
}

// New is
func New(baseURL string) *JPH {
	if baseURL == "" {
		return nil
	}

	return &JPH{
		Posts: NewPosts(baseURL),
	}
}
