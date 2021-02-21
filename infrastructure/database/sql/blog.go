package sql

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

func (br *blogRepository) Fetch() (blogs.Blogs, error) {
	model := blogs.Blogs{}

	if err := br.db.Where("deleted_at IS NULL").Find(&model).Error; err != nil {
		return nil, err
	}

	return model, nil
}

func (br *blogRepository) Create(b *blogs.Blog) error {
	if err := br.db.Create(b).Error; err != nil {
		return err
	}

	return nil
}
