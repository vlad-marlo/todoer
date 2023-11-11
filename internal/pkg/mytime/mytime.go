package mytime

import (
	"fmt"
	"time"
)

const Format = time.RFC3339Nano

type MyTime struct {
	Time time.Time
}

// UnmarshalJSON supports json.Unmarshaler interface
func (t *MyTime) UnmarshalJSON(b []byte) error {
	parsed, err := time.Parse(Format, string(b))
	if err != nil {
		return fmt.Errorf("time: Parse: %w", err)
	}

	t.Time = parsed
	return nil
}

// MarshalJSON supports json.Marshaler interface
func (t *MyTime) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format(Format)), nil
}
