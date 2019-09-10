package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/rmrobinson-textnow/bazel-grpc-example/api/echo/v1"
	"google.golang.org/grpc"
)

func main() {
	var (
		input = flag.String("input", "", "The text to have echoed back at you")
	)
	flag.Parse()

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("127.0.0.1:1337", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := bazel_example_echo_v1.NewEchoServiceClient(conn)
	resp, err := client.Echo(context.Background(), &bazel_example_echo_v1.EchoRequest{
		Phrase: &bazel_example_echo_v1.Phrase{
			Value: *input,
		},
	})

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	fmt.Printf("response: %s\n", resp.Phrase.Value)
}
