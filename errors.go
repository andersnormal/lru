package lru

import (
	"errors"
)

var (
	// ErrNoPositiveSize indicates that the size for the LRU cache
	// must be a positive number.
	ErrNoPositiveSize = errors.New("Must provide a positive size")
)
