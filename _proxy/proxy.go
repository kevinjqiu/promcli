package promql

import (
"log"
"github.com/prometheus/prometheus/storage"
"github.com/prometheus/prometheus/util/testutil"
	"time"
	"golang.org/x/net/context"
)

type TestCommand = testCommand

type t struct {}
func (t t) Fatal(args ...interface{}) {
	log.Fatal(args)
}

func (t t) Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args)
}

func ParseTestCommand(input string) ([]TestCommand, error) {
	t := Test{
		T: t{},
		cmds: []TestCommand{},
	}

	err := t.parse(input)
	if err != nil {
		return t.cmds, err
	}
	return t.cmds, nil
}

func ExecuteTestCommand(tc TestCommand, storage *storage.Storage) (*storage.Storage, error) {
	t := Test{
		T: t{},
		storage: *storage,
		cmds: []TestCommand{tc},
	}
	t.queryEngine = NewEngine(nil, nil, 20, 10*time.Second)
	t.context, t.cancelCtx = context.WithCancel(context.Background())
	return &t.storage, t.exec(tc)
}

func NewTestStorage() *storage.Storage {
	t := Test{
		T: t{},
	}
	storage := testutil.NewStorage(t)
	return &storage
}
