package mytime

import (
	"fmt"
	"time"
)

const Format = time.RFC3339Nano

type Time time.Time

func (t *Time) UnmarshalJSON(b []byte) error {
	parsed, err := time.Parse(Format, string(b))
	if err != nil {
		return fmt.Errorf("time: Parse: %w", err)
	}

	*t = Time(parsed)
	return nil
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte((*time.Time)(t).Format(Format)), nil
}
