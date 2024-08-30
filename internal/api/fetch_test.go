package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"HS/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestFetchCards(t *testing.T) {
	t.Run("Successful fetch", func(t *testing.T) {
		expectedCards := []models.Card{
			{ID: "1", Name: "Card 1"},
			{ID: "2", Name: "Card 2"},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(expectedCards)
		}))
		defer server.Close()

		cards, err := FetchCards(server.URL)

		assert.NoError(t, err)
		assert.Equal(t, expectedCards, cards)
	})

	t.Run("Error in Network", func(t *testing.T) {
		cards, err := FetchCards("http://invalid-url")

		assert.Error(t, err)
		assert.Nil(t, cards)
	})

	t.Run("Invalid JSON response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Invalid JSON"))
		}))
		defer server.Close()

		cards, err := FetchCards(server.URL)

		assert.Error(t, err)
		assert.Nil(t, cards)
	})

	t.Run("Not 200 status code", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		cards, err := FetchCards(server.URL)

		assert.Error(t, err)
		assert.Nil(t, cards)
	})
}

func TestFetchCollectableCards(t *testing.T) {
	t.Run("Successful fetch", func(t *testing.T) {
		expectedCards := []models.CollectableCard{
			{
				Card: models.Card{
					ID:          "EX1_001",
					DbfID:       559,
					Name:        "Leper Gnome",
					Text:        "Deathrattle: Deal 2 damage to the enemy hero.",
					Flavor:      "He really just wants to be your friend, but the constant rejection is starting to really get to him.",
					Artist:      "Glenn Rane",
					Attack:      1,
					CardClass:   "Neutral",
					Collectible: true,
					Cost:        1,
					Elite:       false,
					Faction:     "Neutral",
					Health:      1,
					Mechanics:   []string{"Deathrattle"},
					Rarity:      "Common",
					Set:         "Classic",
					Type:        "Minion",
				},
				Classes:     []string{"Neutral"},
				Race:        "Undead",
				TechLevel:   1,
				SpellDamage: 0,
			},
			{
				Card: models.Card{
					ID:          "EX1_002",
					DbfID:       560,
					Name:        "The Black Knight",
					Text:        "Battlecry: Destroy an enemy minion with Taunt.",
					Flavor:      "He was sent by the Lich King to disrupt the Argent Tournament. We can pretty much mark that a failure.",
					Artist:      "Alex Horley Orlandelli",
					Attack:      4,
					CardClass:   "Neutral",
					Collectible: true,
					Cost:        6,
					Elite:       true,
					Faction:     "Neutral",
					Health:      5,
					Mechanics:   []string{"Battlecry"},
					Rarity:      "Legendary",
					Set:         "Classic",
					Type:        "Minion",
				},
				Classes:         []string{"Neutral"},
				HasDiamondSkin:  true,
				HowToEarnGolden: "Unlocked at Level 60.",
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(expectedCards)
		}))
		defer server.Close()

		cards, err := FetchCollectableCards(server.URL)

		assert.NoError(t, err)
		assert.Equal(t, expectedCards, cards)
	})

	t.Run("Network error", func(t *testing.T) {
		cards, err := FetchCollectableCards("http://invalid-url")

		assert.Error(t, err)
		assert.Nil(t, cards)
	})

	t.Run("Invalid JSON response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Invalid JSON"))
		}))
		defer server.Close()

		cards, err := FetchCollectableCards(server.URL)

		assert.Error(t, err)
		assert.Nil(t, cards)
	})

	t.Run("Not 200 status code", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		cards, err := FetchCollectableCards(server.URL)

		assert.Error(t, err)
		assert.Nil(t, cards)
	})

	t.Run("Empty response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode([]models.CollectableCard{})
		}))
		defer server.Close()

		cards, err := FetchCollectableCards(server.URL)

		assert.NoError(t, err)
		assert.Empty(t, cards)
	})

	t.Run("Partial data", func(t *testing.T) {
		partialCard := models.CollectableCard{
			Card: models.Card{
				ID:   "PARTIAL_001",
				Name: "Partial Card",
			},
		}

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode([]models.CollectableCard{partialCard})
		}))
		defer server.Close()

		cards, err := FetchCollectableCards(server.URL)

		assert.NoError(t, err)
		assert.Len(t, cards, 1)
		assert.Equal(t, partialCard, cards[0])
	})
}
