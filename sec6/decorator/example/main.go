package main

import "github.com/RealWorldGoSolution/sec6/decorator"

func main() {
	if err := decorator.Exec(); err != nil {
		panic(err)
	}
}
