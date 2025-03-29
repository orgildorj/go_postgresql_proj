package scrape

import (
	"os"

	"github.com/joho/godotenv"
)

const coinGeckoBaseURL = "https://api.coingecko.com/api/v3"

func addAPI(url string) string {
	godotenv.Load(".env")
	return url + "&x_cg_demo_api_key=" + os.Getenv("coin_gecko_demo_api")
}
