package pkg

func LivePrefixChanger() (string, bool) {
	var prefix string
	switch ApplicationState.state {
	case StateCommand: prefix = ">>> "
	case StateEval: prefix = "(eval)>>> "
	case StateLoad: prefix = "(load)>>> "
	}
	return prefix, true
}
