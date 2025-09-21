package config

import (
	"log"

	"example.com/blog/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// 初始化数据库连接
func InitDB() {
	var err error
	// 使用 SQLite 作为数据库
	db, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	// 配置数据库连接
	// dsn := "root:guo0909@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 自动迁移模式
	err = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database initialized successfully")
}

// 获取数据库连接
func GetDB() *gorm.DB {
	return db
}
