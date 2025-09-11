package main

import (
	"errors"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint           `gorm:"primaryKey"`                        // 主键
	Name      string         `gorm:"type:varchar(100);not null"`        // 用户名，非空
	Email     string         `gorm:"type:varchar(100);unique;not null"` // 邮箱，唯一且非空
	Age       uint8          // 年龄
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index"` // 软删除字段
}

// TableName 自定义表名
func (User) TableName() string {
	return "user"
}

func createEntity(db *gorm.DB) {
	newUser := User{
		Name:  "王五",
		Email: "wangwu@example.com",
		Age:   35,
	}
	result := db.Create(&newUser)
	if result.Error != nil {
		log.Fatal("Create failed:", result.Error)
	}

	// 创建成功后，newUser 的 ID 等字段会被自动填充
	log.Println("Created user's ID:", newUser.ID)
}

func queryEntity(db *gorm.DB) {
	// 查询单条记录
	var user User
	// 获取第一条记录（主键升序）
	//result := db.First(&user)
	// 根据主键查询
	// result := db.First(&user, 3)
	// 查询指定条件的记录
	result := db.Where("name = ?", "王五").First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("Record not found")
		} else {
			log.Fatal("Query failed:", result.Error)
		}
	}
	log.Printf("Query result: %+v\n", user)

	// 查询多条记录
	var users []User
	result = db.Where("age > ?", 25).Find(&users) // 查询年龄大于 25 的所有用户
	if result.Error != nil {
		log.Fatal("Query failed:", result.Error)
	}

	for _, user := range users {
		log.Printf("User: %+v\n", user)
	}

}

func updateEntity(db *gorm.DB) {

	// 先查询出要更新的记录
	var user User
	// db.First(&user, "name = ?", "张三")
	db.First(&user, 1)

	// 更新单个字段
	db.Model(&user).Update("Age", 31)

	// 更新多个字段（使用结构体，注意零值字段不会被更新）
	db.Model(&user).Updates(User{Name: "张三", Age: 32})

	// 更新多个字段（使用 Map ，可以更新零值）
	db.Model(&user).Updates(map[string]interface{}{"Name": "张三", "Age": 0})

	// 使用 Save 保存所有字段（通常用于更新从数据库查询出的记录）
	user.Email = "zhangsan@example.com"
	db.Save(&user)
}

func deleteEntity(db *gorm.DB) {
	// 软删除 - 模型必须包含 gorm.DeletedAt 字段
	//var user User
	// db.First(&user, 1)
	//db.First(&user, "name = ?", "张三")
	//db.Delete(&user) // 执行软删除，DeleteAt 会被设置为当前时间

	// 永久删除
	// db.Unscoped().Delete(&user)

	// 根据条件删除多条记录
	db.Where("age > ?", 0).Delete(&User{})
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

	dsn := "root:guo0909@tcp(127.0.0.1:3306)/knowledge?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	// 自动迁移 - 根据模型创建或更新表结构
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}

	// db, err := gorm.Open(mysql.New(mysql.Config{
	// 	DSN:                       "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
	// 	DefaultStringSize:         256,                                                                        // string 类型字段的默认长度
	// 	DisableDatetimePrecision:  true,                                                                       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	// 	DontSupportRenameIndex:    true,                                                                       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	// 	DontSupportRenameColumn:   true,                                                                       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	// 	SkipInitializeWithVersion: false,                                                                      // 根据当前 MySQL 版本自动配置
	// }), &gorm.Config{})

	// createEntity(db)

	// queryEntity(db)

	//updateEntity(db)

	deleteEntity(db)

}
