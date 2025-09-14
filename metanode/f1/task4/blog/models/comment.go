package models

import (
	"gorm.io/gorm"
)

// Comment 模型定义评论
type Comment struct {
	gorm.Model
	Content string `gorm:"not null" json:"content"`
	UserID  uint   `gorm:"not null" json:"user_id"`
	User    User   `json:"user,omitempty"`
	PostID  uint   `gorm:"not null" json:"post_id"`
	Post    Post   `json:"post,omitempty"`
}
