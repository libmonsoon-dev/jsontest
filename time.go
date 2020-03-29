package jsontest

import (
	"encoding/json"
	"fmt"
	"time"
)

type Duration time.Duration

var _ json.Marshaler = new(Duration)
var _ json.Unmarshaler = new(Duration)

func newDuration(seconds float64) Duration {
	return Duration(time.Duration(seconds * float64(time.Second)))
}

func (e Duration) MarshalJSON() ([]byte, error) {
	seconds := time.Duration(e).Seconds()
	return json.Marshal(seconds)
}

func (e *Duration) UnmarshalJSON(data []byte) error {
	var seconds float64

	err := json.Unmarshal(data, &seconds)
	if err != nil {
		return fmt.Errorf("could not unmarshal float64: %w", err)
	}

	*e = newDuration(seconds)
	return nil
}
