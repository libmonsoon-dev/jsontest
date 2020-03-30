package jsontest

import (
	"encoding/json"
	"fmt"
	"time"
)

type duration time.Duration

var _ json.Marshaler = new(duration)
var _ json.Unmarshaler = new(duration)

func newDuration(seconds float64) duration {
	return duration(time.Duration(seconds * float64(time.Second)))
}

func (d duration) MarshalJSON() ([]byte, error) {
	seconds := time.Duration(d).Seconds()
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
