package faker_test

import (
	"go-service-gin/util/faker"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSentence(t *testing.T) {
	ac := len(faker.Sentence())

	assert.GreaterOrEqual(t, ac, 25)
}

func TestParagraph(t *testing.T) {
	ac := len(faker.Paragraph())

	assert.GreaterOrEqual(t, ac, 50)
}

// // Paragraph is
// func Paragraph() string {
// 	return v3.Paragraph()
// }
