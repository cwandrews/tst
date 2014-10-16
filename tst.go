package tst

import (
	"flag"
	"fmt"
)

var (
	color    bool
	failures int
)

func init() {
	flag.BoolVar(&color, "color", false, "Colored test output")
}

func Run() int {
	if !flag.Parsed() {
		flag.Parse()
	}
	for _, ts := range allTestSuites {
		if err := ts.Run(); err != nil {
			Fail("Error running %s: %v\n", ts.Name, err)
		}
	}
	if failures > 0 {
		fmt.Printf("There were %d total failures\n", failures)
		return failures
	}
	fmt.Printf("There were no failures in any test suite\n")
	return 0
}
