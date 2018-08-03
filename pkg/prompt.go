package pkg

func LivePrefixChanger() (string, bool) {
	var prefix string
	switch ApplicationState.state {
	case StateCommand:
		prefix = ">>> "
	case StateLoad:
		prefix = "(load)>>> "
	}
	return prefix, true
}
