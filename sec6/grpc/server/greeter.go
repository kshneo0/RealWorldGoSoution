package main

import (
	"fmt"

	"github.com/RealWorldGoSolution/sec6/grpc/greeter"
	"golang.org/x/net/context"
)

type Greeter struct {
	Exclaim bool
}

func (g *Greeter) Greet(ctx context.Context, r *greeter.GreetRequest) (*greeter.GreetResponse, error) {
	msg := fmt.Sprintf("%s %s", r.GetGreeting(), r.GetName())
	if g.Exclaim {
		msg += "!"
	} else {
		msg += "."
	}
	return &greeter.GreetResponse{Response: msg}, nil
}
