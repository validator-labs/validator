package release

import (
	"bytes"
	"time"
)

// emptyString contains an empty JSON string value to be used as output.
var emptyString = `""`

// Time is a convenience wrapper around stdlib time, but with different
// marshalling and unmarshaling for zero values.
type Time struct {
	time.Time
}

// Now returns the current time. It is a convenience wrapper around time.Now().
func Now() Time {
	return Time{time.Now()}
}

// MarshalJSON marshals a Time to JSON.
func (t Time) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte(emptyString), nil
	}
	return t.Time.MarshalJSON()
}

// UnmarshalJSON unmarshals a Time from JSON.
func (t *Time) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		return nil
	}
	// If it is empty, we don't have to set anything since time.Time is not a
	// pointer and will be set to the zero value
	if bytes.Equal([]byte(emptyString), b) {
		return nil
	}
	return t.Time.UnmarshalJSON(b)
}

// Parse parses a formatted string and returns the time value it represents.
func Parse(layout, value string) (Time, error) {
	t, err := time.Parse(layout, value)
	return Time{Time: t}, err
}

// ParseInLocation is like Parse but with a location.
func ParseInLocation(layout, value string, loc *time.Location) (Time, error) {
	t, err := time.ParseInLocation(layout, value, loc)
	return Time{Time: t}, err
}

// Date returns the Time corresponding to the given year, month, day, hour, min, sec, and nsec.
func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) Time {
	return Time{Time: time.Date(year, month, day, hour, min, sec, nsec, loc)}
}

// Unix generates a Time from sec and nsec.
func Unix(sec int64, nsec int64) Time { return Time{Time: time.Unix(sec, nsec)} }

// Add returns the Time t+d.
func (t Time) Add(d time.Duration) Time { return Time{Time: t.Time.Add(d)} }

// AddDate returns the Time corresponding to adding the given number of years, months, and days to t.
func (t Time) AddDate(years int, months int, days int) Time {
	return Time{Time: t.Time.AddDate(years, months, days)}
}

// After reports whether the time instant t is after u.
func (t Time) After(u Time) bool { return t.Time.After(u.Time) }

// Before reports whether the time instant t is before u.
func (t Time) Before(u Time) bool { return t.Time.Before(u.Time) }

// Equal reports whether Time t equals u.
func (t Time) Equal(u Time) bool { return t.Time.Equal(u.Time) }

// In returns a copy of t representing the same time instant, but with a new location.
func (t Time) In(loc *time.Location) Time { return Time{Time: t.Time.In(loc)} }

// Local returns t with the location set to local time.
func (t Time) Local() Time { return Time{Time: t.Time.Local()} }

// Round returns the result of rounding t to the nearest multiple of d (since the zero time).
func (t Time) Round(d time.Duration) Time { return Time{Time: t.Time.Round(d)} }

// Sub returns the duration t-u.
func (t Time) Sub(u Time) time.Duration { return t.Time.Sub(u.Time) }

// Truncate returns t truncated to a multiple of d (since the zero time).
func (t Time) Truncate(d time.Duration) Time { return Time{Time: t.Time.Truncate(d)} }

// UTC returns t with the location set to UTC.
func (t Time) UTC() Time { return Time{Time: t.Time.UTC()} }
