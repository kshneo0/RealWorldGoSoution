package main

import "github.com/RealWorldGoSolution/sec5/pools"

func main() {
	if err := pools.ExecWithTimeout(); err != nil {
		panic(err)
	}
}
