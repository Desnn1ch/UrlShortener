package storage

import "errors"

var (
	ErrURLNotFound = errors.New("not found")
	ErrURLExists   = errors.New("already exists")
)
