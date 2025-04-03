package main

import (
	"fmt"
	"go_backend/internal/scrape"
	"go_backend/internal/sql_db"
	"os"
	"strings"
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
	telegram_ids := strings.Split(os.Getenv("telegram_ids"), ",")
	fmt.Println(today)

	go func() {
		for {
			scrape.ScrapeAusländerbehörde(today, "2025-05-01", telegram_ids, db)
			time.Sleep(30 * time.Second)
		}
	}()

	// r.GET("/ping", test.Ping)
	r.Run("localhost:8081")
}
