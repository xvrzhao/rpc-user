package models

import (
	"errors"
	"github.com/micro-stacks/rpc-user/db"
)

// all 用于数据库迁移，每增删一个 model，这里都需要做相应修改，
// 否则模型变动会同步反应到数据库。
var all = []interface{}{
	new(User),
	new(UserToken),
	new(SmsCodeIP),
}

func Migrate() error {
	if db.Client == nil {
		return errors.New("migrate: db is not connected")
	}

	if err := db.Client.AutoMigrate(all...).Error; err != nil {
		return err
	}
	return nil
}
