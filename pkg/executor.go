package pkg

import (
	"github.com/prometheus/prometheus/storage"
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

func Init() {
	ApplicationState.state = StateCommand
	ApplicationState.buffer = make([]string, 0)
}

func execute(input string, storage *storage.Storage) *storage.Storage {
	cmds, err := promql.ParseTestCommand(input)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, cmd := range cmds {
		storage, err = promql.ExecuteTestCommand(cmd, storage)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return storage
}

func handleClear(input string, storage *storage.Storage) *storage.Storage {
	return execute(input, storage)
}

func handleLoad(input string, storage *storage.Storage) *storage.Storage {
	if ApplicationState.state == StateLoad {
		fmt.Println("Already in Loading state")
		return storage
	}
	ApplicationState.state = StateLoad
	ApplicationState.buffer = append(ApplicationState.buffer, input)
	return storage
}

func handleEval(input string, storage *storage.Storage) *storage.Storage {
	if ApplicationState.state == StateEval {
		fmt.Println("Already in Eval state")
		return storage
	}
	ApplicationState.state = StateEval
	ApplicationState.buffer = append(ApplicationState.buffer, input)
	return storage
}

func handleEmptyLine(input string, storage *storage.Storage) *storage.Storage {
	if ApplicationState.state == StateLoad || ApplicationState.state == StateEval {
		lines := strings.Join(ApplicationState.buffer, "\n")
		storage = execute(lines, storage)
		ApplicationState.buffer = make([]string, 0)
		ApplicationState.state = StateCommand
	}
	return storage
}

func handleOther(input string, storage *storage.Storage) *storage.Storage {
	if ApplicationState.state == StateLoad || ApplicationState.state == StateEval {
		ApplicationState.buffer = append(ApplicationState.buffer, input)
	}
	return storage
}

func Executor(input string, storage *storage.Storage) {
	switch {
	case input == "clear": storage = handleClear(input, storage)
	case strings.HasPrefix(input, "load"): storage = handleLoad(input, storage)
	case strings.HasPrefix(input, "eval"): storage = handleEval(input, storage) // TODO: promql.test's eval evaluates against expectations. We need it to display the result of an evaluation.
	case input == "": storage = handleEmptyLine(input, storage)
	default: storage = handleOther(input, storage)
	}
}

