package jsontest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

const minute = 60
const hour = 60 * minute

func TestEvent_UnmarshalJSON(t *testing.T) {
	defer func(local *time.Location) {
		time.Local = local
	}(time.Local)
	time.Local = time.FixedZone("EEST", 3*hour)

	tests := []struct {
		Arg    []byte
		Expect Event
	}{
		{
			[]byte("{\"Time\":\"2020-03-29T13:54:19.054742+03:00\",\"Action\":\"run\",\"Package\":\"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1\",\"Test\":\"TestAdd\"}"),
			Event{
				time.Date(2020, 3, 29, 13, 54, 19, 54742000, time.Local),
				Run,
				"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1",
				"TestAdd",
				newDuration(0),
				"",
			},
		},
		{
			[]byte("{\"Time\":\"2020-03-29T15:22:34.647704+03:00\",\"Action\":\"output\",\"Package\":\"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1\",\"Test\":\"TestAdd\",\"Output\":\"=== RUN   TestAdd\\n\"}"),
			Event{
				time.Date(2020, 3, 29, 15, 22, 34, 647704000, time.Local),
				Output,
				"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1",
				"TestAdd",
				newDuration(0),
				"=== RUN   TestAdd\n",
			},
		},
		{
			[]byte("{\"Time\":\"2020-03-29T15:22:36.051929+03:00\",\"Action\":\"output\",\"Package\":\"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1\",\"Test\":\"TestAdd\",\"Output\":\"--- PASS: TestAdd (1.40s)\\n\"}"),
			Event{
				time.Date(2020, 3, 29, 15, 22, 36, 51929000, time.Local),
				Output,
				"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1",
				"TestAdd",
				newDuration(0),
				"--- PASS: TestAdd (1.40s)\n",
			},
		},
		{
			[]byte("{\"Time\":\"2020-03-29T15:22:36.052011+03:00\",\"Action\":\"pass\",\"Package\":\"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1\",\"Test\":\"TestAdd\",\"Elapsed\":1.4}"),
			Event{
				time.Date(2020, 3, 29, 15, 22, 36, 52011000, time.Local),
				Pass,
				"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1",
				"TestAdd",
				newDuration(1.4),
				"",
			},
		},
		{
			[]byte("{\"Time\":\"2020-03-29T15:22:36.052031+03:00\",\"Action\":\"output\",\"Package\":\"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1\",\"Output\":\"PASS\\n\"}"),
			Event{
				time.Date(2020, 3, 29, 15, 22, 36, 52031000, time.Local),
				Output,
				"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1",
				"",
				newDuration(0),
				"PASS\n",
			},
		},
		{
			[]byte("{\"Time\":\"2020-03-29T15:22:36.052448+03:00\",\"Action\":\"output\",\"Package\":\"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1\",\"Output\":\"ok  \\tgithub.com/libmonsoon-dev/gomut/testdata/arithmetic/v1\\t1.411s\\n\"}"),
			Event{
				time.Date(2020, 3, 29, 15, 22, 36, 52448000, time.Local),
				Output,
				"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1",
				"",
				newDuration(0),
				"ok  \tgithub.com/libmonsoon-dev/gomut/testdata/arithmetic/v1\t1.411s\n",
			},
		},
		{
			[]byte("{\"Time\":\"2020-03-29T15:22:36.055034+03:00\",\"Action\":\"pass\",\"Package\":\"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1\",\"Elapsed\":1.414}"),
			Event{
				time.Date(2020, 3, 29, 15, 22, 36, 55034000, time.Local),
				Pass,
				"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1",
				"",
				newDuration(1.414),
				"",
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%v", i+1), func(t *testing.T) {
			var event Event
			if err := json.Unmarshal(test.Arg, &event); err != nil {
				t.Errorf("json.Unmarshal() error = %s", err)
			}
			if event != test.Expect {
				t.Errorf("event != test.Expect:\nActual:\t%v\nExpect:\t%v", event, test.Expect)
			}
		})

	}
}

func TestEvent_MarshalJSON(t *testing.T) {
	defer func(local *time.Location) {
		time.Local = local
	}(time.Local)
	time.Local = time.FixedZone("EEST", 3*hour)

	tests := []struct {
		Expect    []byte
		Arg Event
	}{
		{
			[]byte("{\"Time\":\"2020-03-29T13:54:19.054742+03:00\",\"Action\":\"run\",\"Package\":\"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1\",\"Test\":\"TestAdd\"}"),
			Event{
				time.Date(2020, 3, 29, 13, 54, 19, 54742000, time.Local),
				Run,
				"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1",
				"TestAdd",
				newDuration(0),
				"",
			},
		},
		{
			[]byte("{\"Time\":\"2020-03-29T15:22:34.647704+03:00\",\"Action\":\"output\",\"Package\":\"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1\",\"Test\":\"TestAdd\",\"Output\":\"=== RUN   TestAdd\\n\"}"),
			Event{
				time.Date(2020, 3, 29, 15, 22, 34, 647704000, time.Local),
				Output,
				"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1",
				"TestAdd",
				newDuration(0),
				"=== RUN   TestAdd\n",
			},
		},
		{
			[]byte("{\"Time\":\"2020-03-29T15:22:36.051929+03:00\",\"Action\":\"output\",\"Package\":\"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1\",\"Test\":\"TestAdd\",\"Output\":\"--- PASS: TestAdd (1.40s)\\n\"}"),
			Event{
				time.Date(2020, 3, 29, 15, 22, 36, 51929000, time.Local),
				Output,
				"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1",
				"TestAdd",
				newDuration(0),
				"--- PASS: TestAdd (1.40s)\n",
			},
		},
		{
			[]byte("{\"Time\":\"2020-03-29T15:22:36.052011+03:00\",\"Action\":\"pass\",\"Package\":\"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1\",\"Test\":\"TestAdd\",\"Elapsed\":1.4}"),
			Event{
				time.Date(2020, 3, 29, 15, 22, 36, 52011000, time.Local),
				Pass,
				"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1",
				"TestAdd",
				newDuration(1.4),
				"",
			},
		},
		{
			[]byte("{\"Time\":\"2020-03-29T15:22:36.052031+03:00\",\"Action\":\"output\",\"Package\":\"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1\",\"Output\":\"PASS\\n\"}"),
			Event{
				time.Date(2020, 3, 29, 15, 22, 36, 52031000, time.Local),
				Output,
				"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1",
				"",
				newDuration(0),
				"PASS\n",
			},
		},
		{
			[]byte("{\"Time\":\"2020-03-29T15:22:36.052448+03:00\",\"Action\":\"output\",\"Package\":\"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1\",\"Output\":\"ok  \\tgithub.com/libmonsoon-dev/gomut/testdata/arithmetic/v1\\t1.411s\\n\"}"),
			Event{
				time.Date(2020, 3, 29, 15, 22, 36, 52448000, time.Local),
				Output,
				"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1",
				"",
				newDuration(0),
				"ok  \tgithub.com/libmonsoon-dev/gomut/testdata/arithmetic/v1\t1.411s\n",
			},
		},
		{
			[]byte("{\"Time\":\"2020-03-29T15:22:36.055034+03:00\",\"Action\":\"pass\",\"Package\":\"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1\",\"Elapsed\":1.414}"),
			Event{
				time.Date(2020, 3, 29, 15, 22, 36, 55034000, time.Local),
				Pass,
				"github.com/libmonsoon-dev/gomut/testdata/arithmetic/v1",
				"",
				newDuration(1.414),
				"",
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("test #%v", i+1), func(t *testing.T) {
			rawJSON, err := json.Marshal(test.Arg)
			if err != nil {
				t.Errorf("json.Marshal() error = %s", err)
			}
			if !bytes.Equal(rawJSON, test.Expect) {
				t.Errorf("rawJSON != test.Expect:\nActual:\t%s\nExpect:\t%s", rawJSON, test.Expect)
			}
		})

	}
}