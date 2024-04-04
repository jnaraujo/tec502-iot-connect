package types

import "time"

type Time struct {
	time.Time
}

func NewTimeNow() Time {
	return Time{Time: time.Now()}
}

func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		t.Time = time.Time{}
		return nil
	}

	return t.Time.UnmarshalJSON(data)
}

func (t *Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	} else {
		return t.Time.MarshalJSON()
	}
}
