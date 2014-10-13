package tst

import (
	"flag"
)

var color bool

func init() {
	flag.BoolVar(&color, "color", false, "Colored test output")
}

func Run() {
	if !flag.Parsed() {
		flag.Parse()
	}
	for _, ts := range allTestSuites {
		if err := ts.Run(); err != nil {
			Fail("Error running %s: %v\n", ts.Name, err)
		}
	}
}
