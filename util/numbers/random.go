package numbers

import (
	"math/rand"
	"time"
)

// RandomBetween is
func RandomBetween(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// RandomBetweenInt32 is
func RandomBetweenInt32(min int32, max int32) int32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int31n(max-min) + min
}

// RandomBetweenInt64 is
func RandomBetweenInt64(min int64, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min) + min
}

// RandomInt64 is
func RandomInt64() int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63()
}

// RandomInt32 is
func RandomInt32() int32 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int31()
}

// Random is
func Random() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}
