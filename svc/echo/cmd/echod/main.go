package main

import (
	"fmt"
	"log"
	"net"

	"github.com/rmrobinson-textnow/bazel-grpc-example/svc/echo"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 1337))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	bazel_example_echo_v1.RegisterEchoServiceServer(grpcServer, &bazel_example_echo_v1.API{})
	grpcServer.Serve(lis)
}
