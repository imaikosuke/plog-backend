package main

import (
	"log"
	"plog-backend/internal/db"
	"plog-backend/internal/handlers"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// .envファイルの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 環境変数の設定
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		log.Fatal("JWT_SECRET_KEY is not set")
	}

	// データベースの初期化
	db.Init()

	r := gin.Default()

	r.POST("/api/register", handlers.RegisterUser)
	r.POST("/api/login", handlers.LoginUser)

	auth := r.Group("/")
	auth.Use(handlers.AuthMiddleware())
	{
		auth.POST("/api/photolog", handlers.CreatePhotolog)
		auth.GET("/api/photologs", handlers.GetPhotologs)
		auth.DELETE("/api/photolog", handlers.DeletePhotolog)
	}

	r.Run()
}
