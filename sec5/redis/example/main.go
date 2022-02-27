package main

import "github.com/RealWorldGoSolution/sec5/redis"

func main() {
	if err := redis.Exec(); err != nil {
		panic(err)
	}

	if err := redis.Sort(); err != nil {
		panic(err)
	}
}
