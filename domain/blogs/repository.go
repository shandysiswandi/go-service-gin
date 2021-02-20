package blogs

// BlogRepository is
type BlogRepository interface {
	Fetch() ([]Blog, error)
}
