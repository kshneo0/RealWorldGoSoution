package main

import "github.com/RealWorldGoSolution/sec5/storage"

func main() {
	if err := storage.Exec(); err != nil {
		panic(err)
	}
}
