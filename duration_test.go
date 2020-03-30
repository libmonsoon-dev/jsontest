package jsontest

import "encoding/json"

var _ json.Marshaler = new(duration)
var _ json.Unmarshaler = new(duration)
