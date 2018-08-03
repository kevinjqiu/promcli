package main

import (
	"github.com/c-bata/go-prompt"
	"github.com/kevinjqiu/promcli/pkg"
	"fmt"
)

const Banner = `
 ____                       ____ _     ___ 
|  _ \ _ __ ___  _ __ ___  / ___| |   |_ _|
| |_) | '__/ _ \| '_   _ \| |   | |    | | 
|  __/| | | (_) | | | | | | |___| |___ | | 
|_|   |_|  \___/|_| |_| |_|\____|_____|___|`

func main() {
	fmt.Println(Banner)
	pkg.Help("")
	p := prompt.New(
		pkg.Executor,
		pkg.Completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionTitle("PromCLI"),
		prompt.OptionLivePrefix(pkg.LivePrefixChanger),
	)
	p.Run()
}
