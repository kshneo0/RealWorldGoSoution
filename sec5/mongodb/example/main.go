package main

import "github.com/RealWorldGoSolution/sec5/mongodb"

func main() {
	if err := mongodb.Exec(); err != nil {
		panic(err)
	}
}
