package uuid_test

import (
	"go-service-gin/util/uuid"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUUID(t *testing.T) {
	is := assert.New(t)

	cases := []struct {
		input    string
		expected bool
	}{
		{"51e5fe10-5325-4a32-bce8-7ebe9708c453", true},
		{"50ee7e88-c29e-4e38-b94b-6937b5a4c8c ", false},
		{"aeb11973-3be9-47fc-8059-ecdac9ab1887", true},
		{"abcdefgh-ijkl-mnop-qrst-uvwxyzabcdef", false},
		{"688831FF-D2D6-4DA9-8676-AC4C9436BC0E", true},
	}

	for _, tc := range cases {
		is.Equal(tc.expected, uuid.IsUUID(tc.input))
	}
}

func TestGenerate(t *testing.T) {
	is := assert.New(t)

	cases := []struct {
		input    string
		expected bool
	}{
		{uuid.Generate(), true},
		{uuid.Generate(), true},
		{uuid.Generate(), true},
		{uuid.Generate(), true},
		{uuid.Generate(), true},
	}

	for _, tc := range cases {
		is.Equal(tc.expected, uuid.IsUUID(tc.input))
	}
}
