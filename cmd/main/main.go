package main

import (
	"plog-backend/internal/db"
	"plog-backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
    // データベースの初期化
    db.Init()

    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    // ユーザー登録とログインのルーティング
    r.POST("/api/register", handlers.RegisterUser)
    r.POST("/api/login", handlers.LoginUser)

    r.Run()
}
