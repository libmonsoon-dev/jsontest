package jsontest

import (
	"encoding/json"
	"fmt"
	"time"
)

type duration time.Duration

func newDuration(seconds float64) duration {
	return duration(seconds * float64(time.Second))
}

func (d duration) Duration() time.Duration {
	return time.Duration(d)
}

func (d duration) Seconds() float64 {
	return d.Duration().Seconds()
}

func (d duration) MarshalJSON() ([]byte, error) {
	seconds := d.Seconds()
	return json.Marshal(seconds)
}

func (d *duration) UnmarshalJSON(data []byte) error {
	var seconds float64

	err := json.Unmarshal(data, &seconds)
	if err != nil {
		return fmt.Errorf("could not unmarshal float64: %w", err)
	}

	*d = newDuration(seconds)
	return nil
}
