package storage

import "errors"

var (
	ErrTaskDoesNotExists = errors.New("task with provided id does not exists")
)
