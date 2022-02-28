package main

import "github.com/RealWorldGoSolution/sec6/rest"

func main() {
	if err := rest.Exec(); err != nil {
		panic(err)
	}
}
