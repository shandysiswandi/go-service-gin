package application

import "go-service-gin/domain/blogs"

// BlogApplication is
type BlogApplication struct {
	blogRepository blogs.BlogRepository
}

// NewBlogApplication is
func NewBlogApplication(br blogs.BlogRepository) *BlogApplication {
	if br == nil {
		return nil
	}

	return &BlogApplication{br}
}

// Fetch is
func (ba *BlogApplication) Fetch() ([]blogs.Blog, error) {
	return ba.blogRepository.Fetch()
}
