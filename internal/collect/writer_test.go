package collect

import (
	"HS/internal/models"
	"bytes"
	"encoding/json"
	"os"
	"reflect"
	"testing"
)

func TestWriteCardsToCSV(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := os.CreateTemp("", "testcards*.csv")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	cards := []models.Card{ // Sample card data
		{
			ID:          "TEST01",
			DbfID:       123,
			Name:        "Test Card",
			Text:        "This is a test card",
			Flavor:      "Tasty!",
			Artist:      "John Doe",
			Attack:      2,
			CardClass:   "Neutral",
			Collectible: true,
			Cost:        3,
			Elite:       false,
			Faction:     "Horde",
			Health:      4,
			Mechanics:   []string{"Battlecry", "Deathrattle"},
			Rarity:      "Common",
			Set:         "Classic",
			Type:        "Minion",
		},
	}

	err = WriteCardsToCSV(cards, tmpfile.Name()) // call the function
	if err != nil {
		t.Fatalf("WriteCardsToCSV failed: %v", err)
	}

	content, err := os.ReadFile(tmpfile.Name()) // read the created CSV file

	if err != nil {
		t.Fatalf("Failed to read the created CSV file: %v", err)
	}

	// check the content of the CSV file
	expectedContent := `ID,DbfID,Name,Text,Flavor,Artist,Attack,CardClass,Collectible,Cost,Elite,Faction,Health,Mechanics,Rarity,Set,Type
TEST01,123,Test Card,This is a test card,Tasty!,John Doe,2,Neutral,true,3,false,Horde,4,Battlecry|Deathrattle,Common,Classic,Minion
`
	if string(content) != expectedContent {
		t.Errorf("CSV content does not match expected. Got:\n%s\nWant:\n%s", string(content), expectedContent)
	}
}

func TestWriteCardsToJSON(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "testcards*.json") // create a temporary file for testing

	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	cards := []models.CollectableCard{ // Sample card data
		{
			Card: models.Card{
				ID:          "TEST01",
				DbfID:       123,
				Name:        "Test Card",
				Text:        "This is a test card",
				Flavor:      "Tasty!",
				Artist:      "John Doe",
				Attack:      2,
				CardClass:   "Neutral",
				Collectible: true,
				Cost:        3,
				Elite:       false,
				Faction:     "Horde",
				Health:      4,
				Mechanics:   []string{"Battlecry", "Deathrattle"},
				Rarity:      "Common",
				Set:         "Classic",
				Type:        "Minion",
			},
			Classes:                  []string{"Mage", "Warrior"},
			Armor:                    0,
			HasDiamondSkin:           false,
			MercenariesRole:          1,
			HideCost:                 false,
			BattlegroundsBuddyDbfId:  0,
			Race:                     "Beast",
			IsBattlegroundsPoolSpell: false,
			TechLevel:                2,
			HideStats:                false,
			SpellSchool:              "Frost",
		},
	}

	err = WriteCardsToJSON(cards, tmpfile.Name()) // call the function being tested
	if err != nil {
		t.Fatalf("WriteCardsToJSON failed: %v", err)
	}

	content, err := os.ReadFile(tmpfile.Name()) // read the created JSON
	if err != nil {
		t.Fatalf("Failed to read the created JSON file: %v", err)
	}

	var decodedCards []models.CollectableCard // Unmarshal the JSON to verify its struct
	err = json.Unmarshal(content, &decodedCards)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON content: %v", err)
	}

	if !reflect.DeepEqual(cards, decodedCards) {
		t.Errorf("Unmarshaled data does not match original data. Got:\n%+v\nWant:\n%+v", decodedCards, cards)
	}

	var prettyJSON bytes.Buffer // pretty print the JSON for visual inspection
	err = json.Indent(&prettyJSON, content, "", "  ")
	if err != nil {
		t.Fatalf("Failed to pretty print JSON: %v", err)
	}
	t.Logf("Pretty printed JSON:\n%s", prettyJSON.String())
}
