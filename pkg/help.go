package pkg

import "fmt"

const (
	HelpSummary = `
Welcome to PromCLI
==================
This console allows you to test PromQL expressions on user-defined metric fixtures.
The supported commands are:
* load  - load <step:duration>, e.g., load 5m
          Enter the fixture loading state in which you can record one fixture per line.
          An empty line will exit the fixture loading state.
* eval  - eval instant at <instant:duration> <promql_expr>, e.g., eval instant at 0m http_requests{status_code="200"}
          Evaluate PromQL expressions.
          Optionally you can record a set of expectations to evaluate against.
* clear - Clear all fixtures currently in the database and start over.

For detailed explainations and examples, type help <topic>, e.g., help load

Press Ctrl + D to exit`
	HelpLoad = `
load <step:duration>
-----------

Put the shell in the fixture loading state. In this state, you are able to manually populate the fixtures
the Prometheus database will have. You must specify a "step" (in the format of duration, see example below)
argument to indicate how long each metric value will sustain.

e.g.,

    load 5m
        http_requests{status_code="200"} 1

This means that the metric "http_requests" will be set to 1 for 5 minutes.  "5m" is the duration, meaning 5 "minutes".
Other possible unit values are "s" (seconds), "h" (hours), "d" (days).

You are also able to give a metric multiple values, e.g.,

	load 5m
		http_requests{status_code="200"} 1 2 3 4

This will generate the following metric values:

   value
	^
  4 |                   x
  3 |              x----|
  2 |         x----|
  1 |    x----|
	0----|----|----|----|--------> time
		 v    v    v    v
		 5m  10m  15m  20m
"`
	HelpEval = `
eval instant at <instant:duration> <expr>
-----------------------------------------

Evaluate a PromQL expression at "instant" on the current set of metrics fixtures in the database.

"instant" is of the duration form, e.g., 5m (5 minutes), 1s (1 second), 3h (3 hours), 2d (2 days).

e.g.,

	eval instant at 0m http_requests{status_code="200"}
`
	HelpClear = `
clear
-----

Reset the database and remove all metrics fixtures previously loaded.`)

func Help(topic string) {
	switch (topic) {
	case "load": fmt.Println(HelpLoad)
	case "eval": fmt.Println(HelpEval)
	case "clear": fmt.Println(HelpClear)
	default: fmt.Println(HelpSummary)
	}
}

