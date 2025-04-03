package scrape

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go_backend/internal/bot"
	"go_backend/internal/sql_db"
	"log"
	"os"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

func ScrapeAusländerbehörde(startDate string, endDate string, telegram_ids []string, db *sql.DB) {
	scrapeURL := fmt.Sprintf("https://www48.muenchen.de/buergeransicht/api/backend/available-days?startDate=%s&endDate=%s&officeId=10187259&serviceId=10339027&serviceCount=1", startDate, endDate)
	godotenv.Load(".env")

	c := colly.NewCollector()

	c.OnHTML("", func(h *colly.HTMLElement) {
		fmt.Println("Termin verfügbar!")
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting Terminvereinbarung.")
	})

	c.OnResponse(func(r *colly.Response) {
		var data map[string]interface{}

		err := json.Unmarshal(r.Body, &data)
		if err != nil {
			log.Println("Error parsing JSON:", err)
		}

		if code, exists := data["errorCode"].(string); exists && code == "noAppointmentForThisScope" {
			fmt.Println("❌ No appointment available!")
			for _, id := range telegram_ids {
				bot.SendTelegramMessage(id, `❌ No appointment available!`, os.Getenv("telegram_api"))
			}
			sql_db.InsertTerminStatus(db, false)
		} else {
			fmt.Println("✅ Appointment availble!!!")
			for _, id := range telegram_ids {
				bot.SendTelegramMessage(id, `✅ Appointment available!!!
			Go to https://stadt.muenchen.de/buergerservice/terminvereinbarung.html#/services/10339027/locations/10187259`, os.Getenv("telegram_api"))
			}

			sql_db.InsertTerminStatus(db, true)
		}

	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error while scraping: %s", err)
	})

	c.Visit(scrapeURL)
}
