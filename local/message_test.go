package main

import (
	"context"
	"reflect"
	"testing"

	_ "git.code.oa.com/trpc-go/trpc-go/http"
	pb "github.com/Truth-NJU/tprc-go/trpcprotocol/local"
	"github.com/golang/mock/gomock"
)

//go:generate go mod tidy
//go:generate mockgen -destination=stub/github.com/Truth-NJU/tprc-go/trpcprotocol/local/local_mock.go -package=local -self_package=github.com/Truth-NJU/tprc-go/trpcprotocol/local --source=stub/github.com/Truth-NJU/tprc-go/trpcprotocol/local/local.trpc.go

func Test_messageImpl_SendMessage(t *testing.T) {
	// 开始写mock逻辑
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	messageService := pb.NewMockMessageService(ctrl)
	var inorderClient []*gomock.Call
	// 预期行为
	m := messageService.EXPECT().SendMessage(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
		s := &messageImpl{}
		return s.SendMessage(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// 开始写单元测试逻辑
	type args struct {
		ctx context.Context
		req *pb.MessageRequest
		rsp *pb.MessageResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rsp *pb.MessageResponse
			var err error
			if rsp, err = messageService.SendMessage(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("messageImpl.SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("messageImpl.SendMessage() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}
