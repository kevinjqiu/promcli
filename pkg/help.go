package pkg

import "fmt"

func Help() {
	fmt.Println(`
Welcome to PromCLI
==================`)
	fmt.Println(`This console allows you to test PromQL expressions on user-defined metric fixtures.
The supported commands are:
* load  - load <step:duration>, e.g., load 5m
          Enter the fixture loading state in which you can record one fixture per line.
          An empty line will exit the fixture loading state.
* eval  - eval instant at <instant:duration> <promql_expr>, e.g., eval instant at 0m http_requests{status_code="200"}
          Evaluate PromQL expressions.
          Optionally you can record a set of expectations to evaluate against.
* clear - Clear all fixtures currently in the database and start over.`)

	fmt.Println("Press Ctrl+D to exit")
}

