package numbers_test

import (
	"go-service-gin/util/numbers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomInt64(t *testing.T) {
	is := assert.New(t)

	cases := []struct {
		name     string
		expected int64
	}{
		{"Random Int64 0", int64(0)},
		{"Random Int64 1", int64(1)},
		{"Random Int64 2", int64(2)},
		{"Random Int64 3", int64(3)},
		{"Random Int64 4", int64(4)},
		{"Random Int64 5", int64(5)},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			is.NotEqual(tc.expected, numbers.RandomInt64())
		})
	}
}

func TestRandomInt32(t *testing.T) {
	is := assert.New(t)

	cases := []struct {
		name     string
		expected int32
	}{
		{"Random Int32 0", int32(0)},
		{"Random Int32 1", int32(1)},
		{"Random Int32 2", int32(2)},
		{"Random Int32 3", int32(3)},
		{"Random Int32 4", int32(4)},
		{"Random Int32 5", int32(5)},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			is.NotEqual(tc.expected, numbers.RandomInt32())
		})
	}
}

func TestRandom(t *testing.T) {
	is := assert.New(t)

	cases := []struct {
		name     string
		expected int32
	}{
		{"Random Int 0", 0},
		{"Random Int 1", 1},
		{"Random Int 2", 2},
		{"Random Int 3", 3},
		{"Random Int 4", 4},
		{"Random Int 5", 5},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			is.NotEqual(tc.expected, numbers.Random())
		})
	}
}

func TestRandomBetween(t *testing.T) {
	is := assert.New(t)

	cases := []struct {
		name     string
		expected int
	}{
		{"Random Between Int 0", 0},
		{"Random Between Int 1", 1},
		{"Random Between Int 2", 2},
		{"Random Between Int 3", 3},
		{"Random Between Int 4", 4},
		{"Random Between Int 5", 5},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			is.NotEqual(tc.expected, numbers.RandomBetween(6, 10))
		})
	}
}

func TestRandomBetweenInt32(t *testing.T) {
	is := assert.New(t)

	cases := []struct {
		name     string
		expected int32
	}{
		{"Random Between Int32 0", 0},
		{"Random Between Int32 1", 1},
		{"Random Between Int32 2", 2},
		{"Random Between Int32 3", 3},
		{"Random Between Int32 4", 4},
		{"Random Between Int32 5", 5},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			is.NotEqual(tc.expected, numbers.RandomBetweenInt32(6, 10))
		})
	}
}

func TestRandomBetweenInt64(t *testing.T) {
	is := assert.New(t)

	cases := []struct {
		name     string
		expected int32
	}{
		{"Random Between Int64 0", 0},
		{"Random Between Int64 1", 1},
		{"Random Between Int64 2", 2},
		{"Random Between Int64 3", 3},
		{"Random Between Int64 4", 4},
		{"Random Between Int64 5", 5},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			is.NotEqual(tc.expected, numbers.RandomBetweenInt64(6, 10))
		})
	}
}
