package jsontest

import "time"

// https://golang.org/cmd/test2json/#hdr-Output_Format
type Event struct {
	Time    time.Time
	Action  Action
	Package string
	Test    string   `json:",omitempty"`
	Elapsed duration `json:",omitempty"`
	Output  string   `json:",omitempty"`
}
