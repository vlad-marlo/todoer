//go:generate easyjson -all status.go
package model

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Status string

const (
	StatusCreated = "created"
	StatusDeleted = "deleted"
	StatusInWork  = "in-work"
	StatusDone    = "done"
)

var (
	ErrInvalidStatus = errors.New("invalid status")
)

func statusValid(s Status) error {
	switch s {
	case StatusDone, StatusCreated, StatusDeleted, StatusInWork, "":
		return nil
	}
	return fmt.Errorf("%w: %s", ErrInvalidStatus, s)
}

func (s *Status) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*s = Status(raw)
	return statusValid(*s)
}
