package api

import (
	"HS/internal/models"
	"encoding/json"
	"io"
	"net/http"
)

func FetchCards(apiURL string) ([]models.Card, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cards []models.Card
	err = json.Unmarshal(body, &cards)
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func FetchCollectableCards(apiURL string) ([]models.CollectableCard, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cards []models.CollectableCard
	err = json.Unmarshal(body, &cards)
	if err != nil {
		return nil, err
	}

	return cards, nil
}
