package models

import (
	"gorm.io/gorm"
)

// Post 模型定义博客文章
type Post struct {
	gorm.Model
	Title    string    `gorm:"not null" json:"title"`
	Content  string    `gorm:"not null" json:"content"`
	UserID   uint      `gorm:"not null" json:"user_id"`
	User     User      `json:"user,omitempty"`
	Comments []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty"`
}
