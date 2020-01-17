package models

import (
	"github.com/jinzhu/gorm"
	"github.com/micro-stacks/rpc-user/db"
	pwd "github.com/micro-stacks/utils/password"
)

// User 对应用户表 users。
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);not null;unique_index;comment:'用户名'"`
	Nickname string `gorm:"type:varchar(50);not null;comment:'用户昵称'"`
	Gender   uint8  `gorm:"type:tinyint;default:0;comment:'性别（0 未知 1 男 2 女）'"`
	Phone    string `gorm:"type:varchar(20);unique_index;comment:'手机号'"`
	Email    string `gorm:"type:varchar(80);default:null;unique_index;comment:'邮箱'"`
	PwdHash  string `gorm:"type:char(32);not null;comment:'密码哈希'"`
	Salt     string `gorm:"type:char(8);not null;comment:'盐'"`
}

// hasUsername 判断是否存在用户名。
func (m User) HasUsername(username string) (has bool, err error) {
	var count int
	if err = db.Client.Model(&m).Where("username = ?", username).Count(&count).Error; err != nil {
		return
	}
	if count > 0 {
		has = true
	} else {
		has = false
	}
	return
}

// HasPhone 判断手机号是否已经注册。
func (m User) HasPhone(phone string) (has bool, err error) {
	var count int
	if err = db.Client.Model(&m).Where("phone = ?", phone).Count(&count).Error; err != nil {
		return
	}
	if count > 0 {
		has = true
	} else {
		has = false
	}
	return
}

// Register 注册用户。
func (m User) Register(username, nickname, phone, password string) (userID int64, err error) {
	shp, s := pwd.SaltHashPwd(password, 8)
	m.Username = username
	m.Nickname = nickname
	m.Phone = phone
	m.PwdHash = shp
	m.Salt = s
	if err = db.Client.Create(&m).Error; err != nil {
		return
	}
	userID = int64(m.ID)
	return
}
