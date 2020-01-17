package models

import (
	"fmt"
	"log"
)

// 运行本包内的单元测试都会先执行迁移
func init() {
	fmt.Println("executing migration...")
	if err := Migrate(); err != nil {
		log.Fatalf("failed to migrate: %v", err)
		return
	}
}
