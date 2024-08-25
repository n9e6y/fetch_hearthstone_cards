package main

import (
	"HS/internal/api"
	"HS/internal/csv"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env")
	}

	apiURL := os.Getenv("API_URL")
	if apiURL == "" {
		log.Fatal("API_URL not set in env")
	}

	cards, err := api.FetchCards(apiURL)
	if err != nil {
		log.Fatalf("Error fetching cards: %v", err)
	}

	err = csv.WriteCardsToCSV(cards, "cards.csv")
	if err != nil {
		log.Fatalf("Error writing cards to CSV: %v", err)
	}

	fmt.Printf("Successfully wrote %d cards to cards.csv\n", len(cards))
}
