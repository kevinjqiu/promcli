package main

import (
	"github.com/c-bata/go-prompt"
	"github.com/kevinjqiu/promcli/pkg"
	"fmt"
	"github.com/prometheus/prometheus/promql"
)

const Banner = `
 ____                       ____ _     ___ 
|  _ \ _ __ ___  _ __ ___  / ___| |   |_ _|
| |_) | '__/ _ \| '_   _ \| |   | |    | | 
|  __/| | | (_) | | | | | | |___| |___ | | 
|_|   |_|  \___/|_| |_| |_|\____|_____|___|`


func executor(input string) {
	storage := promql.NewTestStorage()
	pkg.Executor(input, storage)
}

func main() {
	fmt.Println(Banner)
	p := prompt.New(
		executor,
		pkg.Completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionTitle("PromCLI"),
		prompt.OptionLivePrefix(pkg.LivePrefixChanger),
	)
	p.Run()
}
