package pkg

import (
		"strings"
	"fmt"
	"github.com/prometheus/prometheus/promql"
)

const (
	StateCommand = iota
	StateLoad
	StateEval
)

var ApplicationState struct {
	state int
	buffer []string
}

var testShell *promql.TestShell

func init() {
	ApplicationState.state = StateCommand
	ApplicationState.buffer = make([]string, 0)
	testShell = promql.NewTestShell()
}

func execute(input string) {
	cmds, err := promql.ParseTestCommand(input)
	if err != nil {
		fmt.Println(err.Error())
	}

	testShell.SetCmds(cmds)
	err = testShell.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func handleClear(input string) {
	execute(input)
}

func handleLoad(input string) {
	if ApplicationState.state == StateLoad {
		fmt.Println("Already in Loading state")
		return
	}
	ApplicationState.state = StateLoad
	ApplicationState.buffer = append(ApplicationState.buffer, input)
	return
}

func handleEval(input string) {
	if ApplicationState.state == StateEval {
		fmt.Println("Already in Eval state")
		return
	}
	ApplicationState.state = StateEval
	ApplicationState.buffer = append(ApplicationState.buffer, input)
	return
}

func handleEmptyLine(input string) {
	if ApplicationState.state == StateLoad || ApplicationState.state == StateEval {
		lines := strings.Join(ApplicationState.buffer, "\n")
		execute(lines)
		ApplicationState.buffer = make([]string, 0)
		ApplicationState.state = StateCommand
	}
}

func handleOther(input string) {
	if ApplicationState.state == StateLoad || ApplicationState.state == StateEval {
		ApplicationState.buffer = append(ApplicationState.buffer, input)
	}
}

func Executor(input string) {
	switch {
	case input == "clear": handleClear(input)
	case strings.HasPrefix(input, "load"): handleLoad(input)
	case strings.HasPrefix(input, "eval"): handleEval(input)
	case input == "": handleEmptyLine(input)
	default: handleOther(input)
	}
}

