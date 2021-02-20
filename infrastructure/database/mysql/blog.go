package mysql

import (
	"go-service-gin/domain/blogs"
	"go-service-gin/infrastructure/database"

	"gorm.io/gorm"
)

type blogRepository struct {
	db *gorm.DB
}

// NewBlogRepository is
func NewBlogRepository(db *database.Database) blogs.BlogRepository {
	if db == nil {
		return nil
	}

	return &blogRepository{db.SQL}
}

func (br *blogRepository) Fetch() ([]blogs.Blog, error) {
	model := []blogs.Blog{}

	if err := br.db.Find(&model).Error; err != nil {
		return nil, err
	}

	return model, nil
}
