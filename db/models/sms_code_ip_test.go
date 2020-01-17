package models

import (
	"github.com/micro-stacks/rpc-user/db"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestSmsCodeIP_IpAddOne(t *testing.T) {
	m := new(SmsCodeIP)
	ip := "test" + strconv.FormatInt(time.Now().UnixNano(), 10) // 每次单元测试确保 IP 唯一
	// 没有当天该 IP 的记录时
	if err := m.IpAddOne(ip); err != nil {
		t.Error(0, err)
		return
	}
	m1 := new(SmsCodeIP)
	err := db.Client.Where("ip = ? and created_at > ?", ip, time.Now().Format("2006-01-02")).First(m1).Error
	if err != nil {
		t.Error(1, err)
		return
	}
	if m1.Num != 1 {
		t.Error(2, m1.Num)
		return
	}
	// 有当天该 IP 的记录时
	if err := m.IpAddOne(ip); err != nil {
		t.Error(3, err)
		return
	}
	m2 := new(SmsCodeIP)
	err = db.Client.Where("ip = ? and created_at > ?", ip, time.Now().Format("2006-01-02")).First(m2).Error
	if err != nil {
		t.Error(4, err)
		return
	}
	if m2.Num != 2 {
		t.Error(5)
		return
	}
}

func TestSmsCodeIP_CheckIP(t *testing.T) {
	ip := strconv.FormatInt(time.Now().UnixNano(), 10)
	m := new(SmsCodeIP)
	max, err := strconv.Atoi(os.Getenv("SMS_CODE_MAX_REQ_PER_IP"))
	if err != nil {
		t.Error(0, err)
		return
	}
	valid, err := m.CheckIP(ip)
	if err != nil {
		t.Error(1, err)
		return
	}
	if valid == false {
		t.Error(2)
		return
	}
	for i := 0; i < max+1; i++ {
		err = m.IpAddOne(ip)
		if err != nil {
			t.Error(3)
			return
		}
	}
	valid, err = m.CheckIP(ip)
	if err != nil {
		t.Error(4)
		return
	}
	if valid == true {
		t.Error(5)
	}
}
