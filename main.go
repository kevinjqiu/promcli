package main

import (
	"github.com/c-bata/go-prompt"
	"fmt"
	"github.com/prometheus/prometheus/promql"
)

func executor(input string) {
	storage := promql.NewTestStorage()
	fmt.Println("Your input: " + input)
	cmds, err := promql.ParseTestCommand(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(cmds)
	for _, cmd := range cmds {
		storage, err = promql.ExecuteTestCommand(cmd, storage)
		if err != nil {
			panic(err)
		}
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "users", Description: "Users table"},
		{Text: "articles", Description: "Articles table"},
		{Text: "comments", Description: "Comments table"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionTitle("PromCLI"),
	)

	p.Run()
}
