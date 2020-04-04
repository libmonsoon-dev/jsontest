package jsontest

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os/exec"
)

func ExampleEvent() {
	type message struct {
		event Event
		error error
	}

	runTests := func(packages ...string) <-chan message {
		ch := make(chan message)

		go func() {
			defer close(ch)
			if len(packages) == 0 {
				packages = []string{"./..."}
			}
			cmd := exec.Command("go", append([]string{"test", "-json"}, packages...)...)
			stdout, err := cmd.StdoutPipe()

			if err != nil {
				ch <- message{error: fmt.Errorf("could not pipe test cmd stdout: %w", err)}
				return
			}

			if err := cmd.Start(); err != nil {
				ch <- message{error: fmt.Errorf("could not start test cmd: %w", err)}
				return
			}

			r := bufio.NewReader(stdout)
			var event Event

			for {
				line, err := r.ReadBytes('\n')

				if err != nil {
					if errors.Is(err, io.EOF) {
						break
					}
					ch <- message{error: fmt.Errorf("could not read cmd stdout: %w", err)}
					return
				}

				if err := json.Unmarshal(line, &event); err != nil {
					ch <- message{error: fmt.Errorf("could not unmarshal test event \"%s\": %w", line, err)}
					return
				}
				ch <- message{event: event}
			}

			if err := cmd.Wait(); err != nil {
				ch <- message{error: fmt.Errorf("test command error: %w", err)}
			}
		}()

		return ch
	}

	for msg := range runTests("bufio", "encoding/json", "errors", "fmt", "io", "os/exec") {
		if msg.error != nil {
			fmt.Println(msg.error)
		} else {
			fmt.Printf("%#v\n", msg.event)
		}
	}
}
