package tst

import (
	"fmt"
	"time"
)

type Test struct {
	test_func    func(t *TestSuite) error
	fatal        bool
	name         string
	parent_suite *TestSuite
}

func (t *Test) Run() {
	start := time.Now()
	err := t.test_func(t.parent_suite)
	dur := time.Since(start).Seconds()
	if err != nil && t.fatal {
		panic(fmt.Errorf("TST FATAL ERROR: %v\n", err))
	} else if err != nil {
		t.parent_suite.fail("Test %s FAIL\n%v\nDuration: %.3f sec\n", t.name, err, dur)
	} else {
		t.parent_suite.pass("Test %s PASS\nDuration: %.3f sec\n", t.name, dur)
	}
}
