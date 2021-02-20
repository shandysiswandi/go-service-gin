package uuid

import (
	"regexp"

	UUID "github.com/google/uuid"
)

// IsUUID is
func IsUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

// Generate is
func Generate() string {
	return UUID.NewString()
}
