package jsontest

type Action string

const (
	Run    Action = "run"    // the test has started running
	Pause         = "pause"  // the test has been paused
	Cont          = "cont"   // the test has continued running
	Pass          = "pass"   // the test passed
	Bench         = "bench"  // the benchmark printed log output but did not fail
	Fail          = "fail"   // the test or benchmark failed
	Output        = "output" // the test printed output
	Skip          = "skip"   // the test was skipped or the package contained no tests
)