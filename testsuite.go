package tst

import (
	"fmt"
)

var allTestSuites []*TestSuite

type TestSuite struct {
	Name       string
	setup      func(*TestSuite) error
	teardown   func(*TestSuite) error
	tests      []*Test
	vars       map[string]interface{}
	num_passed int
	num_failed int
	num_tests  int
}

func DefTestSuite(name string) *TestSuite {
	ts := &TestSuite{
		Name:  name,
		tests: []*Test{},
		vars:  map[string]interface{}{},
	}
	ts.setup = func(ts *TestSuite) error { return nil }
	ts.teardown = func(ts *TestSuite) error { return nil }
	allTestSuites = append(allTestSuites, ts)
	return ts
}

func (tt *TestSuite) SetSetup(setup_func func(*TestSuite) error) {
	tt.setup = setup_func
}

func (tt *TestSuite) SetTearDown(teardown_func func(*TestSuite) error) {
	tt.teardown = teardown_func
}

func (tt *TestSuite) DefTest(name string, test_func func(*TestSuite) error, fatal bool) {
	t := &Test{
		name:         name,
		test_func:    test_func,
		fatal:        fatal,
		parent_suite: tt,
	}
	tt.num_tests++
	tt.tests = append(tt.tests, t)
}

func (tt *TestSuite) Run() error {
	if setup_err := tt.setup(tt); setup_err != nil {
		return fmt.Errorf("Error during %s setup: %v\n", tt.Name, setup_err)
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("RECOVER: %+v\n", r)
		}
	}()
	for _, t := range tt.tests {
		t.Run()
	}
	if teardown_err := tt.teardown(tt); teardown_err != nil {
		return fmt.Errorf("Error durring %s teardown: %v\n", tt.Name, teardown_err)
	}
	tt.PrintResults()
	return nil
}

func (tt *TestSuite) PrintResults() {
	if tt.num_passed == tt.num_tests {
		Pass("=========================================\n100%% of %s tests passed\n%d passes\n%d failures\n", tt.Name, tt.num_passed, tt.num_failed)
	} else {
		Fail("=========================================\n%.2f%% of %s tests passed\n%d passes\n%d failures\n", (float64(tt.num_passed)/float64(tt.num_passed))*100, tt.Name, tt.num_passed, tt.num_failed)
	}
}

func (tt *TestSuite) pass(s string, stuff ...interface{}) {
	tt.num_passed++
	Pass(s, stuff...)
}

func (tt *TestSuite) fail(s string, stuff ...interface{}) {
	tt.num_failed++
	Fail(s, stuff...)
}

func (tt *TestSuite) GetVar(key string) interface{} {
	return tt.vars[key]
}

func (tt *TestSuite) SetVar(key string, val interface{}) {
	tt.vars[key] = val
}

func (tt *TestSuite) Varp(key string) bool {
	_, ok := tt.vars[key]
	return ok
}
