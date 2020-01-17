package models

import (
	"github.com/jinzhu/gorm"
	"github.com/micro-stacks/rpc-user/db"
	"os"
	"strconv"
	"time"
)

// SmsIP 对应 sms_code_ips 表，表示某个 IP 在当天请求发送短信验证码的次数。
type SmsCodeIP struct {
	gorm.Model
	IP  string `gorm:"type:varchar(100);not null;index;comment:'请求发送验证码的 IP'"`
	Num int64  `gorm:"type:int;not null;default:1;comment:'请求发送验证码的次数'"`
}

// IpAddOne 添加当天的 ip 的请求发送次数
func (m SmsCodeIP) IpAddOne(ip string) error {
	err := db.Client.Where("ip = ? and created_at > ?", ip, time.Now().Format("2006-01-02")).First(&m).Error
	if err == gorm.ErrRecordNotFound {
		m.IP = ip
		m.Num = 1
		if err = db.Client.Create(&m).Error; err != nil {
			return err
		}
		return nil
	}
	if err != nil {
		return err
	}
	m.Num++
	return db.Client.Save(&m).Error
}

// CheckIP 检查 IP 是否超过请求发送限制。
func (m SmsCodeIP) CheckIP(ip string) (valid bool, err error) {
	err = db.Client.Where("ip = ? and created_at > ?", ip, time.Now().Format("2006-01-02")).First(&m).Error
	if err == gorm.ErrRecordNotFound {
		valid = true
		err = nil
		return
	}
	if err != nil {
		return
	}
	max, err := strconv.Atoi(os.Getenv("SMS_CODE_MAX_REQ_PER_IP"))
	if err != nil {
		return
	}
	if m.Num > int64(max) {
		return
	}
	valid = true
	return
}
