package main

import (
	"net/http"
	"plog-backend/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
		db.Init()
    r := gin.Default()

    // Pingテスト用のエンドポイント
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })
		

    // 他のエンドポイントを追加
    r.POST("/api/upload", uploadHandler)

    // サーバーの起動
    r.Run() // デフォルトでは ":8080" でリッスンします
}

// 画像アップロードのハンドラー
func uploadHandler(c *gin.Context) {
    // アップロード処理をここに記述
    c.JSON(http.StatusOK, gin.H{"message": "Upload endpoint hit"})
}
