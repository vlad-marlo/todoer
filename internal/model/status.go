//go:generate easyjson -all status.go
package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Status string

func (s *Status) String() string {
	return string(*s)
}

const (
	StatusCreated = "created"
	StatusDeleted = "deleted"
	StatusInWork  = "in-work"
	StatusDone    = "done"
)

var (
	ErrInvalidStatus = errors.New("invalid status (valid are: created, deleted, in-work, done)")
)

func statusValid(s Status) error {
	switch s {
	case StatusDone, StatusCreated, StatusDeleted, StatusInWork, "":
		return nil
	}
	return fmt.Errorf("%w: %s", ErrInvalidStatus, s)
}

func ParseStatus(raw string) (*Status, error) {
	s := Status(raw)
	err := statusValid(s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func ParseManyStatuses(raw, separator string) (res []Status, err error) {
	var s *Status
	for _, status := range strings.Split(raw, separator) {
		s, err = ParseStatus(status)
		if err != nil {
			return nil, fmt.Errorf("%w: got %s", ErrInvalidStatus, status)
		}
		res = append(res, *s)
	}
	return
}

func (s *Status) SetStatusWhichDoesNotContainsInStatuses(statuses ...Status) {
	statusMap := map[string]struct{}{
		StatusInWork:  {},
		StatusDeleted: {},
		StatusCreated: {},
		StatusDone:    {},
	}
	for _, status := range statuses {
		if _, ok := statusMap[status.String()]; ok {
			delete(statusMap, status.String())
		}
	}
	var count int
	for status := range statusMap {
		count++
		*s = Status(status)
	}
	if count != 1 {
		panic(fmt.Sprintf("unexpected count of statuses: %v", statusMap))
	}
}

func (s *Status) Int() int {
	switch *s {
	case StatusCreated:
		return 1
	case StatusInWork:
		return 2
	case StatusDone:
		return 3
	case StatusDeleted:
		return 4
	}
	return 0
}

func StatusFromInt(status int) Status {
	switch status {
	case 1:
		return StatusCreated
	case 2:
		return StatusInWork
	case 3:
		return StatusDone
	case 4:
		return StatusDeleted
	}
	return StatusCreated
}

func (s *Status) UnmarshalJSON(data []byte) error {
	var raw string

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	*s = Status(raw)
	return statusValid(*s)
}
