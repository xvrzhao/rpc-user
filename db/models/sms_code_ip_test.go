package models

import (
	"github.com/micro-stacks/rpc-user/db"
	"strconv"
	"testing"
	"time"
)

func TestSmsCodeIP_IpAddOne(t *testing.T) {
	m := new(SmsCodeIP)
	ip := "test" + strconv.FormatInt(time.Now().UnixNano(), 10) // 每次单元测试确保 IP 唯一

	// 没有当天该 IP 的记录时
	if err := m.IpAddOne(ip); err != nil {
		t.Error(err)
		return
	}
	err := db.Client.Where("ip = ? and created_at > ?", ip, time.Now().Format("2006-01-02")).First(m).Error
	if err != nil {
		t.Error(err)
		return
	}
	if m.Num != 1 {
		t.Error("not 1")
		return
	}

	// 有当天该 IP 的记录时
	if err := m.IpAddOne(ip); err != nil {
		t.Error(err)
		return
	}
	err = db.Client.Where("ip = ? and created_at > ?", ip, time.Now().Format("2006-01-02")).First(m).Error
	if err != nil {
		t.Error(err)
		return
	}
	if m.Num != 2 {
		t.Error("not 2")
		return
	}
}
