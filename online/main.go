// Package main 是由 trpc-go-cmdline v2.4.0 生成的服务端示例代码
// 注意：本文件并非必须存在，而仅为示例，用户应按需进行修改使用，如不需要，可直接删去

package main

import (
	_ "git.code.oa.com/trpc-go/trpc-filter/debuglog"
	_ "git.code.oa.com/trpc-go/trpc-filter/recovery"
	trpc "git.code.oa.com/trpc-go/trpc-go"
	"git.code.oa.com/trpc-go/trpc-go/log"
	pb "github.com/Truth-NJU/tprc-go/trpcprotocol/online"
)

func main() {
	s := trpc.NewServer()

	pb.RegisterMessageService(s, &messageImpl{})
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
