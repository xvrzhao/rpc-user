package server

import (
	"context"
	"github.com/micro-stacks/rpc-user/proto"
)

func (*UserServer) RegisterUser(context.Context, *proto.RegisterUserReq) (*proto.RegisterUserRsp, error) {
	panic("implement me")
}
