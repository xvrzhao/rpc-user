package server

import (
	"code.aliyun.com/microstack/rpc-user/proto"
	"context"
)

func (*UserServer) RegisterUser(context.Context, *proto.RegisterUserReq) (*proto.RegisterUserRsp, error) {
	panic("implement me")
}
