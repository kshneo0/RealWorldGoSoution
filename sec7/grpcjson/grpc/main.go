package main

import (
	"fmt"
	"net"

	"github.com/RealWorldGoSolution/sec7/grpcjson/internal"
	"github.com/RealWorldGoSolution/sec7/grpcjson/keyvalue"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	keyvalue.RegisterKeyValueServer(grpcServer, internal.NewKeyValue())
	lis, err := net.Listen("tcp", ":4444")
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port :4444")
	grpcServer.Serve(lis)
}
