package main

import (
	"fmt"

	"github.com/RealWorldGoSolution/sec4/basicerrors"
)

func main() {
	basicerrors.BasicErrors()

	err := basicerrors.SomeFunc()
	fmt.Println("custom error: ", err)
}
