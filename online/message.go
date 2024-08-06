package main

import (
	"context"

	pb "github.com/Truth-NJU/tprc-go/trpcprotocol/online"
)

type messageImpl struct {
	pb.UnimplementedMessage
}

func (s *messageImpl) SendMessage(
	ctx context.Context,
	req *pb.MessageRequest,
) (*pb.MessageResponse, error) {
	rsp := &pb.MessageResponse{}
	return rsp, nil
}
