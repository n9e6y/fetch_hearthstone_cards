package json

import (
	"HS/internal/models"
	"encoding/json"
	"os"
)

func WriteExpandedCardsToJSON(cards []models.ExpandedCard, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(cards)
}
