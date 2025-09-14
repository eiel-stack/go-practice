package models

import (
	"gorm.io/gorm"
)

// User 模型定义用户信息
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"` // 密码不在JSON中显示
	Email    string `gorm:"unique;not null" json:"email"`
	Posts    []Post `gorm:"foreignKey:UserID" json:"posts,omitempty"`
}
