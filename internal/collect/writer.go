package collect

import (
	"HS/internal/models"
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

func WriteCardsToCSV(cards []models.Card, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"ID", "DbfID", "Name", "Text", "Flavor", "Artist", "Attack", "CardClass", "Collectible", "Cost", "Elite", "Faction", "Health", "Mechanics", "Rarity", "Set", "Type"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, card := range cards {
		row := []string{
			card.ID,
			strconv.Itoa(card.DbfID),
			card.Name,
			card.Text,
			card.Flavor,
			card.Artist,
			strconv.Itoa(card.Attack),
			card.CardClass,
			strconv.FormatBool(card.Collectible),
			strconv.Itoa(card.Cost),
			strconv.FormatBool(card.Elite),
			card.Faction,
			strconv.Itoa(card.Health),
			strings.Join(card.Mechanics, "|"),
			card.Rarity,
			card.Set,
			card.Type,
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

func WriteCardsToJSON(cards []models.CollectableCard, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(cards)
}
