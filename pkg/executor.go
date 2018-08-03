package pkg

import (
	"fmt"
	"github.com/prometheus/prometheus/promql"
	"strings"
)

const (
	StateCommand = iota
	StateLoad
)

var ApplicationState struct {
	state  int
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
	execute(input)
}

func handleEmptyLine(_ string) {
	if ApplicationState.state == StateLoad {
		lines := strings.Join(ApplicationState.buffer, "\n")
		execute(lines)
		ApplicationState.buffer = make([]string, 0)
		ApplicationState.state = StateCommand
	}
}

func handleOther(input string) {
	if ApplicationState.state == StateLoad {
		ApplicationState.buffer = append(ApplicationState.buffer, input)
	}
}

func handleHelp(input string) {
	tokens := strings.Split(input, " ")
	if len(tokens) == 1 {
		Help("")
	} else {
		Help(tokens[1])
	}
}

func Executor(input string) {
	switch {
	case strings.HasPrefix(input, "help"):
		handleHelp(input)
	case input == "clear":
		handleClear(input)
	case strings.HasPrefix(input, "load"):
		handleLoad(input)
	case strings.HasPrefix(input, "eval"):
		handleEval(input)
	case input == "":
		handleEmptyLine(input)
	default:
		handleOther(input)
	}
}
