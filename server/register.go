package server

import (
	"context"
	"github.com/micro-stacks/rpc-user/cache"
	"github.com/micro-stacks/rpc-user/db/models"
	"github.com/micro-stacks/rpc-user/proto"
	"github.com/micro-stacks/utils/protobuf"
)

func (*UserServer) RegisterUser(ctx context.Context, req *proto.RegisterUserReq) (rsp *proto.RegisterUserRsp, err error) {
	rsp = new(proto.RegisterUserRsp)
	if req.GetUserName() == "" ||
		req.GetNickName() == "" ||
		req.GetPassword() == "" ||
		req.GetPhone() == "" ||
		req.GetCode() == "" {
		rsp.Status = protobuf.FalseValue()
		rsp.Msg = "存在空值"
		return
	}
	// 检测手机和验证码是否匹配
	valid, err := cache.VerifyMobileCode(req.GetPhone(), req.GetCode())
	if err != nil {
		return
	}
	if valid == false {
		rsp.Status = protobuf.FalseValue()
		rsp.Msg = "验证码错误"
		return
	}
	// 检测 username 是否已注册
	m := new(models.User)
	has, err := m.HasUsername(req.GetUserName())
	if err != nil {
		return
	}
	if has == true {
		rsp.Status = protobuf.FalseValue()
		rsp.Msg = "用户名已被注册"
		return
	}
	// 检测 phone 是否已注册
	has, err = m.HasPhone(req.GetPhone())
	if err != nil {
		return
	}
	if has == true {
		rsp.Status = protobuf.FalseValue()
		rsp.Msg = "该手机号已被其他账号绑定"
		return
	}
	// 注册用户
	uid, err := m.Register(req.GetUserName(), req.GetNickName(), req.GetPhone(), req.GetPassword())
	if err != nil {
		return
	}
	rsp.Status = protobuf.TrueValue()
	rsp.UserID = uid
	rsp.Msg = "注册成功"
	return
}
