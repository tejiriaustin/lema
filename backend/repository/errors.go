package repository

import "errors"

var (
	ErrConcurrentModification = errors.New("concurrent modification detected")

	ErrNotFound = errors.New("not found")
)
