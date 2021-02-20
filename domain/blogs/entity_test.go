package blogs_test

import (
	"go-service-gin/domain/blogs"
	"testing"
)

func TestTableName(t *testing.T) {
	ex := "blogs"
	ac := blogs.Blog{}.TableName()

	if ac != ex {
		t.Errorf("Expected %s, but actual %s", ex, ac)
	}
}
