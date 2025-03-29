package main

import (
	"fmt"
	"go_backend/internal/scrape"
	"go_backend/internal/sql_db"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	godotenv.Load(".env")
	db := sql_db.InitDB()
	defer db.Close()

	today := string(time.Now().Format("2006-01-02"))
	fmt.Println(today)

	go func() {
		for {
			scrape.ScrapeAusländerbehörde(today, "2025-05-01", os.Getenv("telegram_id"), os.Getenv("telegram_api"), db)
			time.Sleep(60 * time.Second)
		}
	}()

	// r.GET("/ping", test.Ping)
	r.Run("localhost:8081")
}
