package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// User 对应用户表 users。
type User struct {
	gorm.Model
	Username       string     `gorm:"type:varchar(50);not null;unique_index;comment:'用户名'"`
	Nickname       string     `gorm:"type:varchar(50);not null;comment:'用户昵称'"`
	Gender         uint8      `gorm:"type:tinyint;default:0;comment:'性别（0 未知 1 男 2 女）'"`
	Phone          string     `gorm:"type:varchar(20);unique_index;comment:'手机号'"`
	Email          string     `gorm:"type:varchar(80);unique_index;comment:'邮箱'"`
	PwdHash        string     `gorm:"type:char(32);not null;comment:'密码哈希'"`
	Salt           string     `gorm:"type:char(32);not null;comment:'盐'"`
	LastCode       string     `gorm:"type:varchar(20);comment:'最后验证码'"`
	LastCodeExpire *time.Time `gorm:"type:timestamp;comment:'最后验证码过期时间'"`
}
