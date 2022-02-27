package main

import (
	"fmt"

	"github.com/RealWorldGoSolution/sec4/errwrap"
)

func main() {
	errwrap.Wrap()
	fmt.Println()
	errwrap.Unwrap()
	fmt.Println()
	errwrap.StackTrace()
}
