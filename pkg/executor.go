package pkg

import (
	"fmt"
	"github.com/prometheus/prometheus/promql"
)

func Executor(input string) {
	storage := promql.NewTestStorage()
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
}

