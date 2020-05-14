package server

import (
	"context"
	"github.com/micro-stacks/rpc-user/cache"
	"github.com/micro-stacks/rpc-user/db/models"
	"github.com/micro-stacks/rpc-user/proto"
	"github.com/micro-stacks/utils/strings"
)

func (server *UserServer) SendMobileCode(ctx context.Context, req *proto.SendMobileCodeReq) (rsp *proto.SendMobileCodeRsp, err error) {
	// 添加 IP 请求发送次数
	m := new(models.SmsCodeIP)
	err = m.IpAddOne(req.GetIP())
	if err != nil {
		return
	}
	// 检测 IP
	valid, err := m.CheckIP(req.GetIP())
	if err != nil {
		return
	}
	rsp = new(proto.SendMobileCodeRsp)
	if valid == false {
		rsp.Status = false
		rsp.Msg = "IP 超过发送限制"
		return
	}
	// 检测 mobile
	valid, err = cache.CheckMobile(req.GetMobileNumber())
	if err != nil {
		return
	}
	if valid == false {
		rsp.Status = false
		rsp.Msg = "手机号请求验证码过于频繁"
		return
	}
	// 发送验证码
	code := strings.RandNum(6)
	if err = sendMsg(req.GetMobileNumber(), code); err != nil {
		rsp.Status = false
		rsp.Msg = "验证码发送失败"
		return
	}
	// 将验证码放入缓存
	if err = cache.StoreMobileCode(req.GetMobileNumber(), code); err != nil {
		rsp.Status = false
		rsp.Msg = "验证码已发送，但加入缓存失败"
		return
	}
	rsp.Status = true
	rsp.Msg = "发送成功"
	return
}

func sendMsg(mobile, msg string) error {
	// TODO: implementation
	return nil
}
