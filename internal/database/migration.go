package database

import (
	"github.com/zagss/todo-list/internal/database/model"
	"log"
)

// 执行数据迁移
func migration() {
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{})
	if err != nil {
		log.Fatalf("database.migration err: %v", err)
	}
}
