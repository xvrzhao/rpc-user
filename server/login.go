package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/micro-stacks/rpc-user/proto"
)

func (*UserServer) LoginPwd(context.Context, *proto.LoginPwdReq) (*proto.LoginRsp, error) {
	panic("implement me")
}

func (*UserServer) LoginPhone(context.Context, *proto.LoginPhoneReq) (*proto.LoginRsp, error) {
	panic("implement me")
}

func (*UserServer) Logout(context.Context, *empty.Empty) (*wrappers.BoolValue, error) {
	panic("implement me")
}

func (server *UserServer) Auth(context.Context, *proto.AuthReq) (*wrappers.BoolValue, error) {
	panic("implement me")
}
