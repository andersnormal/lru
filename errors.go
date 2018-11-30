package lru

import (
	"errors"
)

var (
	ErrNoPositiveSize = errors.New("Must provide a positive size")
)
