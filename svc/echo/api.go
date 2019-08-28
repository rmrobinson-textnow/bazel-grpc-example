package bazel_example_echo_v1

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

)

type API struct {

}

func (a *API) Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, status.New(codes.InvalidArgument, "bad data").Err()
	}

	req.Phrase.ProcessingTime = ptypes.TimestampNow()

	return &EchoResponse{
		Phrase: req.Phrase,
	}, nil
}