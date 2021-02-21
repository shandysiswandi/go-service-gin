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
	result := blogs.Blogs{}

	if err := br.db.Table(blogs.Blog{}.TableName()).Where("deleted_at IS NULL").Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (br *blogRepository) Create(b *blogs.CreateBlog) error {
	if err := br.db.Table(blogs.Blog{}.TableName()).Create(b).Error; err != nil {
		return err
	}

	return nil
}
