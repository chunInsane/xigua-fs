package main

import (
	"log"
	"xigua-server/config"
	"xigua-server/handlers"
	"xigua-server/middleware"
	"xigua-server/models"
	"xigua-server/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg := config.New()

	if err := utils.EnsureDir("./storage"); err != nil {
		log.Fatal("Failed to create storage directory:", err)
	}

	if err := utils.EnsureDir(cfg.StoragePath); err != nil {
		log.Fatal("Failed to create files storage directory:", err)
	}

	db, err := gorm.Open(sqlite.Open(cfg.DatabasePath), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.AutoMigrate(&models.User{}, &models.FileInfo{}, &models.Folder{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000", "http://localhost:8888", "http://127.0.0.1:8888", "http://10.23.2.18:8888"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	authHandler := handlers.NewAuthHandler(db, cfg)
	fileHandler := handlers.NewFileHandler(db, cfg)
	folderHandler := handlers.NewFolderHandler(db, cfg)
	adminHandler := handlers.NewAdminHandler(db, cfg)

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.GET("/me", middleware.AuthMiddleware(cfg.JWTSecret), authHandler.Me)
		}

		files := api.Group("/files")
		files.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		{
			files.POST("/upload", fileHandler.Upload)
			files.GET("/", fileHandler.List)
			files.GET("/:id/download", fileHandler.Download)
			files.GET("/:id/preview", fileHandler.Preview)
			files.PUT("/:id/move", fileHandler.Move)
			files.DELETE("/:id", fileHandler.Delete)
		}

		folders := api.Group("/folders")
		folders.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		{
			folders.POST("/", folderHandler.Create)
			folders.GET("/", folderHandler.List)
			folders.PUT("/:id", folderHandler.Update)
			folders.PUT("/:id/move", folderHandler.Move)
			folders.DELETE("/:id", folderHandler.Delete)
			folders.GET("/:id/path", folderHandler.GetPath)
		}

		admin := api.Group("/admin")
		admin.Use(middleware.AuthMiddleware(cfg.JWTSecret), middleware.AdminMiddleware())
		{
			admin.GET("/users", adminHandler.GetUsers)
			admin.PUT("/users/:id", adminHandler.UpdateUser)
			admin.DELETE("/users/:id", adminHandler.DeleteUser)
			admin.GET("/stats", adminHandler.GetSystemStats)
		}
	}

	r.Static("/static", "./web/dist")
	r.NoRoute(func(c *gin.Context) {
		c.File("./web/dist/index.html")
	})

	log.Printf("Server starting on port %s", cfg.Port)
	log.Fatal(r.Run(":" + cfg.Port))
}