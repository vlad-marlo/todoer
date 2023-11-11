package mytime

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestTime_UnmarshalJSON(t *testing.T) {
	tt := []time.Time{
		time.Date(2000, time.April, 1, 2, 32, 35, 91, time.UTC),
		time.Date(2020, time.April, 22, 1, 2, 35, 21, time.UTC),
		time.Date(2010, time.September, 14, 16, 10, 12, 52, time.UTC),
		time.Date(2012, time.August, 26, 12, 0, 3, 3, time.UTC),
	}
	for _, tc := range tt {
		t.Run("test time", func(t *testing.T) {
			raw := tc.Format(time.RFC3339Nano)

			got := new(MyTime)

			require.NoError(t, got.UnmarshalJSON([]byte(raw)))
			assert.Equal(t, tc, got.Time)
		})
	}
}

func TestTime_MarshalJSON(t *testing.T) {
	tt := []time.Time{
		time.Date(2000, time.April, 1, 2, 32, 35, 91, time.UTC),
		time.Date(2020, time.April, 22, 1, 2, 35, 21, time.UTC),
		time.Date(2010, time.September, 14, 16, 10, 12, 52, time.UTC),
		time.Date(2012, time.August, 26, 12, 0, 3, 3, time.UTC),
	}
	for _, tc := range tt {
		t.Run("test time", func(t *testing.T) {
			raw := tc.Format(time.RFC3339Nano)

			got := &MyTime{Time: tc}

			marshalled, err := got.MarshalJSON()
			require.NoError(t, err)
			assert.Equal(t, raw, string(marshalled))
		})
	}
}
