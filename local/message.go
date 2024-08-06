package main

import (
	"context"
	pb "github.com/Truth-NJU/tprc-go/trpcprotocol/local"
	"github.com/google/uuid"
)

type messageImpl struct {
	pb.UnimplementedMessage
}

// SendMessage 返回消息id
func (s *messageImpl) SendMessage(
	ctx context.Context,
	req *pb.MessageRequest,
) (*pb.MessageResponse, error) {
	newUUID := uuid.New().String() + "_" + req.SendFrom
	rsp := &pb.MessageResponse{
		MsgId: newUUID,
	}
	return rsp, nil
}
