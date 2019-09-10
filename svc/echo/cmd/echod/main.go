package main

import (
	"fmt"
	"log"
	"net"

	"github.com/rmrobinson-textnow/bazel-grpc-example/api/echo/v1"
	"github.com/rmrobinson-textnow/bazel-grpc-example/svc/echo"
	"google.golang.org/grpc"
)

const (
	port = 1337
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	bazel_example_echo_v1.RegisterEchoServiceServer(grpcServer, &echo.API{})

	log.Printf("Listening on %d\n", port)
	grpcServer.Serve(lis)
}
