package database

import (
	"fmt"
	"github.com/zagss/todo-list/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func Setup() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DatabaseSetting.User,
		conf.DatabaseSetting.Password,
		conf.DatabaseSetting.Host,
		conf.DatabaseSetting.Name)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("database.Setup err: %v", err)
	}

	DB = db
	_db, _ := db.DB()
	_db.SetMaxIdleConns(20)  // 设置连接池，空闲
	_db.SetMaxOpenConns(100) // 打开
	_db.SetConnMaxLifetime(time.Second * 30)

	migration()
}
