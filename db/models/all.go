package models

// All 用于数据库迁移，每增删一个 model，这里都需要做相应修改，
// 否则模型变动会同步反应到数据库。
var All = []interface{}{
	new(User),
	new(UserToken),
}
