package main

import "github.com/RealWorldGoSolution/sec4/global"

func main() {
	if err := global.UseLog(); err != nil {
		panic(err)
	}
}
