package api

import (
	"HS/internal/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func FetchExpandedCards(apiURL string) ([]models.ExpandedCard, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var cards []models.ExpandedCard
	err = json.Unmarshal(body, &cards)
	if err != nil {
		return nil, err
	}

	return cards, nil
}
