package promql

import (
"log"
	"fmt"
	"time"
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

type TestShell struct {
	Test Test
}

func (ts *TestShell) SetCmds(cmds []TestCommand) {
	ts.Test.cmds = cmds
}

func (ts *TestShell) Run() error {
	for _, cmd := range ts.Test.cmds {
		err := ts.exec(cmd)
		if err != nil {
			return err  // TODO: use multi-error
		}
	}
	return nil
}

func (ts *TestShell) exec(tc testCommand) error {
	switch cmd := tc.(type) {
	case *evalCmd:
		t := ts.Test
		q, _ := t.queryEngine.NewInstantQuery(t.storage, cmd.expr, cmd.start)
		res := q.Exec(t.context)
		if res.Err != nil {
			if cmd.fail {
				return nil
			}
			return fmt.Errorf("error evaluating query %q: %s", cmd.expr, res.Err)
		}
		defer q.Close()
		if res.Err == nil && cmd.fail {
			return fmt.Errorf("expected error evaluating query but got none")
		}

		if len(cmd.expected) == 0 {
			fmt.Println(res.Value)
			return nil
		}
		err := cmd.compareResult(res.Value)
		if err != nil {
			return fmt.Errorf("error in %s %s: %s", cmd, cmd.expr, err)
		}

		// Check query returns same result in range mode,
		/// by checking against the middle step.
		q, _ = t.queryEngine.NewRangeQuery(t.storage, cmd.expr, cmd.start.Add(-time.Minute), cmd.start.Add(time.Minute), time.Minute)
		rangeRes := q.Exec(t.context)
		if rangeRes.Err != nil {
			return fmt.Errorf("error evaluating query %q in range mode: %s", cmd.expr, rangeRes.Err)
		}
		defer q.Close()
		if cmd.ordered {
			// Ordering isn't defined for range queries.
			return nil
		}
		mat := rangeRes.Value.(Matrix)
		vec := make(Vector, 0, len(mat))
		for _, series := range mat {
			for _, point := range series.Points {
				if point.T == timeMilliseconds(cmd.start) {
					vec = append(vec, Sample{Metric: series.Metric, Point: point})
					break
				}
			}
		}
		if len(cmd.expected) == 0 {
			fmt.Println(res.Value)
			return nil
		}
		if _, ok := res.Value.(Scalar); ok {
			err = cmd.compareResult(Scalar{V: vec[0].Point.V})
		} else {
			err = cmd.compareResult(vec)
		}
		if err != nil {
			return fmt.Errorf("error in %s %s rande mode: %s", cmd, cmd.expr, err)
		}
	default: return ts.Test.exec(tc)
	}
	return nil
}

func NewTestShell() *TestShell {
	ts := &TestShell{
		Test: Test{
			T: t{},
			cmds: []TestCommand{},
		}}
	ts.Test.exec(&clearCmd{})
	return ts
}
