package main

import (
	"HS/internal/api"
	"HS/internal/collect"
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

	cardsURL := os.Getenv("CARDS_API")
	colcardsURL := os.Getenv("COLLECTABLE_CARDS_API")
	if cardsURL == "" || colcardsURL == "" {
		log.Fatal("vars not in env file")
	}

	//Fetch and collect regular cards
	cards, err := api.FetchCards(cardsURL)
	if err != nil {
		log.Fatalf("Error fetching cards: %v", err)
	}

	err = collect.WriteCardsToCSV(cards, "cards.csv")
	if err != nil {
		log.Fatalf("Error writing cards to CSV: %v", err)
	}

	fmt.Printf("Successfully wrote %d cards to cards.save\n", len(cards))

	// Fetch and collect collectable cards
	expandedCards, err := api.FetchCollectableCards(colcardsURL)
	if err != nil {
		log.Fatalf("Error fetching expanded cards: %v", err)
	}

	err = collect.WriteCardsToJSON(expandedCards, "cards.json")
	if err != nil {
		log.Fatalf("Error writing cards to JSON: %v", err)
	}
}
