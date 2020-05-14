package server

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/micro-stacks/rpc-user/proto"
)

func (*UserServer) LoginPwd(ctx context.Context, req *proto.LoginPwdReq) (rsp *proto.LoginRsp, err error) {
	// 验证账号密码是否正确

	// 生成 token

	panic("")
}

func (*UserServer) LoginPhone(ctx context.Context, req *proto.LoginPhoneReq) (rsp *proto.LoginRsp, err error) {
	// 判断用户库中是否存在该手机

	// 验证码是否正确

	panic("implement me")
}

func (*UserServer) Logout(ctx context.Context, req *empty.Empty) (rsp *wrappers.BoolValue, err error) {
	panic("implement me")
}

func (server *UserServer) Auth(ctx context.Context, req *proto.AuthReq) (rsp *wrappers.BoolValue, err error) {
	panic("implement me")
}
