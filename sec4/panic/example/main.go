package main

import (
	"fmt"

	"github.com/RealWorldGoSolution/sec4/panic"
)

func main() {
	fmt.Println("before panic")
	panic.Catcher()
	fmt.Println("after panic")
}
