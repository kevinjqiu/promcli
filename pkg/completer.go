package pkg

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/c-bata/go-prompt"
)

func completeDuration(d prompt.Document) []prompt.Suggest {
	var s []prompt.Suggest
	w := d.GetWordBeforeCursorWithSpace()
	num, err := strconv.Atoi(w)
	if err != nil {
		s = []prompt.Suggest{
			{Text: "1m", Description: "1 minute"},
			{Text: "5m", Description: "5 minutes"},
			{Text: "1h", Description: "1 hour"},
			{Text: "5h", Description: "5 hours"},
		}
	} else {
		for _, unit := range []string{"m", "h", "d"} {
			s = append(s, prompt.Suggest{
				Text:        fmt.Sprintf("%d%s", num, unit),
				Description: fmt.Sprintf("%d%s", num, unit),
			})
		}
	}

	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func completeEval(d prompt.Document) []prompt.Suggest {
	return completeDuration(d)
}

func completeLoad(d prompt.Document) []prompt.Suggest {
	return completeDuration(d)
}

func completeMetrics(d prompt.Document) []prompt.Suggest {
	// TODO: complete against the metrics in the current session
	return []prompt.Suggest{}
}

func completeExpectations(d prompt.Document) []prompt.Suggest {
	// TODO: complete against the metrics in the current session
	return []prompt.Suggest{}
}

func Completer(d prompt.Document) []prompt.Suggest {
	// w := d.GetWordBeforeCursorWithSpace()
	w := d.TextBeforeCursor()

	switch ApplicationState.state {
	case StateCommand:
		switch {
		case strings.HasPrefix(w, "eval instant at"):
			return completeEval(d)
		case strings.HasPrefix(w, "load"):
			return completeLoad(d)
		default:
			s := []prompt.Suggest{
				{Text: "clear", Description: "Clear the database"},
				{Text: "load", Description: "Enter the fixture loading state.  load <step:duration>"},
				{Text: "eval instant at", Description: "Evaluate expressions.  eval instant at <step:duration> <metric>"},
				//{Text: "eval_fail", Description: "Evaluate expressions"},
				//{Text: "eval_ordered", Description: "Evaluate expressions"},
			}
			return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
		}
		return []prompt.Suggest{}
	case StateEval:
		return completeExpectations(d)
	case StateLoad:
		return completeMetrics(d)
	}
	return []prompt.Suggest{}
}
