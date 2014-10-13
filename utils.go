package tst

import (
	"fmt"
)

const (
	RED   = "\x1b[31;1m"
	GREEN = "\x1b[32;1m"
	NONE  = "\x1b[0m"
)

func Pass(s string, stuff ...interface{}) {
	if color {
		fmt.Printf(Green(s), stuff...)
	} else {
		fmt.Printf(s, stuff...)
	}
}

func Fail(s string, stuff ...interface{}) {
	if color {
		fmt.Printf(Red(s), stuff...)
	} else {
		fmt.Printf(s, stuff...)
	}
}

func Green(s string) string {
	return GREEN + s + NONE
}

func Red(s string) string {
	return RED + s + NONE
}
