package stringy_test

import (
	"go-service-gin/util/stringy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnakeCase(t *testing.T) {
	is := assert.New(t)

	cases := []struct {
		desc     string
		input    string
		expected string
	}{
		{"All text is capitalize only one word", "CAPITALIZE", "capitalize"},
		{"All text is lower with one word", "lower", "lower"},
		{"Text input is IsUUID", "IsUUID", "is_uuid"},
		{"Text input is UserId", "UserId", "user_id"},
		{"Text input is theURL", "theURL", "the_url"},
	}

	for _, tc := range cases {
		is.Equal(tc.expected, stringy.SnakeCase(tc.input))
	}
}
