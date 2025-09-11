package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// User 用户模型
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Posts     []Post `gorm:"foreignKey:UserID"`
	PostCount int    `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Post 文章模型
type Post struct {
	ID            uint      `gorm:"primaryKey"`
	Title         string    `gorm:"size:200;not null"`
	Content       string    `gorm:"type:text;not null"`
	UserID        uint      `gorm:"not null"`
	User          User      `gorm:"foreignKey:UserID"`
	Comments      []Comment `gorm:"foreignKey:PostID"`
	CommentCount  int       `gorm:"default:0"`
	CommentStatus string    `gorm:"size:20;default:'有评论'"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Comment 评论模型
type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	Content   string `gorm:"type:text;not null"`
	UserID    uint   `gorm:"not null"`
	PostID    uint   `gorm:"not null"`
	User      User   `gorm:"foreignKey:UserID"`
	Post      Post   `gorm:"foreignKey:PostID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func getUserPostsWithComments(db *gorm.DB, userID uint) ([]Post, error) {
	var posts []Post
	err := db.Where("user_id = ?", userID).
		Preload("Comments").
		Preload("Comments.User").
		Find(&posts).Error
	return posts, err
}

func getMostCommentedPostAdvanced(db *gorm.DB) (Post, error) {
	var post Post
	err := db.
		Select("posts.*, COUNT(comments.id) as comment_count").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count DESC").
		First(&post).Error
	return post, err
}

// 在文章创建后更新用户的文章数量统计
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	// 更新用户的文章数量
	result := tx.Model(&User{}).
		Where("id = ?", p.UserID).
		Update("post_count", gorm.Expr("post_count + 1"))

	return result.Error
}

// 在评论删除后检查文章评论状态
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	// 查询当前文章的评论数量
	var commentCount int64
	tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&commentCount)

	// 如果评论数量为0，更新文章状态
	if commentCount == 0 {
		result := tx.Model(&Post{}).
			Where("id = ?", c.PostID).
			Updates(map[string]interface{}{
				"comment_status": "无评论",
				"comment_count":  0,
			})
		return result.Error
	}

	// 更新评论数量
	result := tx.Model(&Post{}).
		Where("id = ?", c.PostID).
		Update("comment_count", gorm.Expr("comment_count - 1"))

	return result.Error
}

func main() {
	// 配置数据库连接
	dsn := "root:guo0909@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 自动迁移创建表
	// err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	// if err != nil {
	// 	log.Fatal("数据库迁移失败:", err)
	// }

	// fmt.Println("博客系统数据库初始化完成！")

	//查询用户的所有文章及评论
	posts, err := getUserPostsWithComments(db, 1)
	if err != nil {
		log.Fatal("查询用户文章及评论失败:", err)
	}

	// 打印返回结果
	fmt.Println("用户的文章及评论列表：")
	for _, post := range posts {
		fmt.Printf("文章标题: %s, 评论数量: %d\n", post.Title, len(post.Comments))
		for _, comment := range post.Comments {
			fmt.Printf("评论内容: %s, 评论用户: %s\n", comment.Content, comment.User.Name)
		}
	}

	//查询评论数量最多的文章
	mostCommentedPost, err := getMostCommentedPostAdvanced(db)
	if err != nil {
		log.Fatal("查询评论数量最多的文章失败:", err)
	}

	// 打印返回结果
	fmt.Printf("评论数量最多的文章: %s, 评论数量: %d\n", mostCommentedPost.Title, mostCommentedPost.CommentCount)

}
