package jsontest

import "time"

// https://golang.org/cmd/test2json/#hdr-Output_Format
type Event struct {
	Time    time.Time `json:",omitempty"`
	Action  Action    ``
	Package string    `json:",omitempty"`
	Test    string    `json:",omitempty"`
	Elapsed duration  `json:",omitempty"`
	Output  string    `json:",omitempty"`
}
