package main

import (
	"log"
	"time"

	"example.com/blog/config"
	"example.com/blog/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化数据库连接
	config.InitDB()

	// 设置 Gin 模式，默认是 debug 模式
	gin.SetMode(gin.ReleaseMode)

	// 创建 Gin 引擎
	r := gin.Default()

	// 配置 CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 注册路由
	registerRoutes(r)

	// 启动服务器
	log.Println("Server is running on port :8083")
	r.Run(":8083")
}

// 注册所有路由
func registerRoutes(r *gin.Engine) {
	// 公开路由
	public := r.Group("/api")
	{
		// 用户相关
		public.POST("/register", service.Register)
		public.POST("/login", service.Login)

		// 文章相关（不需要认证的）
		public.GET("/posts", service.GetAllPosts)
		public.GET("/posts/:id", service.GetPostByID)
		public.GET("/posts/:id/comments", service.GetCommentsByPostID)
	}

	// 需要认证的路由
	protected := r.Group("/api")
	protected.Use(service.AuthMiddleware())
	{
		// 文章相关（需要认证的）
		protected.POST("/posts", service.CreatePost)
		protected.PUT("/posts/:id", service.UpdatePost)
		protected.DELETE("/posts/:id", service.DeletePost)

		// 评论相关
		protected.POST("/posts/:id/comments", service.CreateComment)
	}
}
