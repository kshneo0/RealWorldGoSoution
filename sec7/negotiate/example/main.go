package main

import (
	"fmt"
	"net/http"

	"github.com/RealWorldGoSolution/sec7/negotiate"
)

func main() {
	http.HandleFunc("/", negotiate.Handler)
	fmt.Println("Listening on port :3333")
	err := http.ListenAndServe(":3333", nil)
	panic(err)
}
