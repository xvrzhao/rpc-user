package models

import (
	g "github.com/micro-stacks/utils/db/gorm"
	"time"
)

// UserToken 对应用户令牌表 user_tokens。
type UserToken struct {
	g.Model
	UserID      uint64     `gorm:"type:int;not null;unique_index;comment:'用户ID'"`
	Token       string     `gorm:"type:char(32);not null;comment:'用户令牌'"`
	TokenExpire *time.Time `gorm:"type:timestamp;not null;default:current_timestamp;comment:'令牌过期时间'"`
}
