package main

import "github.com/RealWorldGoSolution/sec1/tempfiles"

func main() {
	if err := tempfiles.WorkWithTemp(); err != nil {
		panic(err)
	}
}
