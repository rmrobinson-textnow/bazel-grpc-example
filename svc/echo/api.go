package echo

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/rmrobinson-textnow/bazel-grpc-example/api/echo/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

)

type API struct {
}

func (a *API) Echo(ctx context.Context, req *bazel_example_echo_v1.EchoRequest) (*bazel_example_echo_v1.EchoResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "bad data").Err()
	}

	req.Phrase.ProcessingTime = ptypes.TimestampNow()

	return &bazel_example_echo_v1.EchoResponse{
		Phrase: req.Phrase,
	}, nil
}