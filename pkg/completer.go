package pkg

import "github.com/c-bata/go-prompt"

func Completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "clear", Description: "Clear the database"},
		{Text: "load", Description: "Enter the fixture loading state"},
		{Text: "eval", Description: "Evaluate expressions"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

