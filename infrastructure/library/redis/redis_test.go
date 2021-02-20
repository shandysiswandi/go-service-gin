package redis_test

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	os.Setenv("REDIS_ADDRESS", "localhost:1234")
	os.Setenv("REDIS_PASSWORD", "password")
	os.Setenv("REDIS_DATABASE", "0")
}

func TestNew(t *testing.T) {
	is := assert.New(t)

	log.Println(is)
}

func TestNew_Get(t *testing.T) {
	is := assert.New(t)

	log.Println(is)
}

func TestNew_Set(t *testing.T) {
	is := assert.New(t)

	log.Println(is)
}

func TestNew_Increment(t *testing.T) {
	is := assert.New(t)

	log.Println(is)
}

func TestNew_Decrement(t *testing.T) {
	is := assert.New(t)

	log.Println(is)
}
