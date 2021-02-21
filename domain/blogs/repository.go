package blogs

// BlogRepository is
type BlogRepository interface {
	Fetch() (Blogs, error)
	Create(*Blog) error
}
