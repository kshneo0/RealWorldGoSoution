package main

import (
	"fmt"

	"github.com/RealWorldGoSolution/sec4/log"
)

func main() {
	fmt.Println("basic logging and modification of logger:")
	log.Log()
	fmt.Println("logging 'handled' errors:")
	log.FinalDestination()
}
