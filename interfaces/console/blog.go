package console

import (
	"go-service-gin/application"
	"go-service-gin/domain/blogs"
	"go-service-gin/util/faker"
	"go-service-gin/util/uuid"
)

// BlogREST is
type BlogREST struct {
	blogLogic *application.BlogApplication
}

// NewBlogREST is
func NewBlogREST(ba *application.BlogApplication) *BlogREST {
	if ba == nil {
		return nil
	}

	return &BlogREST{ba}
}

// Create is
func (bh *BlogREST) Create(run int) bool {
	b := &blogs.Blog{}

	for i := 1; i <= run; i++ {
		b = &blogs.Blog{
			ID:    uuid.Generate(),
			Title: faker.Sentence(),
			Body:  faker.Paragraph(),
		}

		if err := bh.blogLogic.Create(b); err != nil {
			return false
		}
	}

	return true
}
