package models

import (
	"github.com/micro-stacks/rpc-user/db"
	"strconv"
	"testing"
	"time"
)

func TestUser_HasUsername(t *testing.T) {
	unique := strconv.FormatInt(time.Now().UnixNano(), 10)
	m := new(User)

	// 没有该用户的情况
	has, err := m.HasUsername(unique)
	if err != nil {
		t.Error(0, err)
		return
	}
	if has == true {
		t.Error(1)
		return
	}

	// 有该用户的情况
	user := &User{
		Username: unique,
		Nickname: "xavier",
		Phone:    unique,
		Email:    unique,
	}
	if err = db.Client.Create(&user).Error; err != nil {
		t.Error(2, err)
		return
	}
	has, err = m.HasUsername(unique)
	if err != nil {
		t.Error(3, err)
		return
	}
	if has == false {
		t.Error(4)
	}
}

func TestUser_HasPhone(t *testing.T) {
	uni := strconv.FormatInt(time.Now().UnixNano(), 10)
	m := new(User)
	has, err := m.HasPhone(uni)
	if err != nil {
		t.Error(0, err)
		return
	}
	if has == true {
		t.Error(1)
		return
	}

	u := &User{
		Username: uni,
		Nickname: uni,
		Phone:    uni,
		Email:    uni,
	}
	if err = db.Client.Create(&u).Error; err != nil {
		t.Error(2, err)
		return
	}
	has, err = m.HasPhone(uni)
	if err != nil {
		t.Error(3, err)
		return
	}
	if has == false {
		t.Error(4)
	}
}

func TestUser_Register(t *testing.T) {
	uni := strconv.FormatInt(time.Now().UnixNano(), 10)
	m := new(User)
	userID, err := m.Register(uni, "register test", uni, "123")
	if err != nil {
		t.Error(1, err)
		return
	}
	db.Client.Where("id = ?", userID).Take(&m)
	if int64(m.ID) != userID || m.Phone != uni || m.Username != uni || m.Nickname != "register test" {
		t.Errorf("2 %#v", m)
	}
}
