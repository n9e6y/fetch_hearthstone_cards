package main

import (
	"HS/internal/api"
	"HS/internal/csv"
	"HS/internal/json"
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
	expandedAPIURL := os.Getenv("EXPANDED_API_URL")
	if apiURL == "" || expandedAPIURL == "" {
		log.Fatal("API_URL or EXPANDED_API_URL not set in .env file")
	}

	//if expandedAPIURL == "" {
	//	log.Fatalf("error")
	//}

	//Fetch and save regular cards
	cards, err := api.FetchCards(apiURL)
	if err != nil {
		log.Fatalf("Error fetching cards: %v", err)
	}

	err = csv.WriteCardsToCSV(cards, "cards.csv")
	if err != nil {
		log.Fatalf("Error writing cards to CSV: %v", err)
	}

	fmt.Printf("Successfully wrote %d cards to cards.csv\n", len(cards))

	// Fetch and save expanded cards
	expandedCards, err := api.FetchExpandedCards(expandedAPIURL)
	if err != nil {
		log.Fatalf("Error fetching expanded cards: %v", err)
	}

	err = json.WriteExpandedCardsToJSON(expandedCards, "expanded_cards.json")
	if err != nil {
		log.Fatalf("Error writing expanded cards to JSON: %v", err)
	}
}
