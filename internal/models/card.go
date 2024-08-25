package models

type Card struct {
	ID          string   `json:"id"`
	DbfID       int      `json:"dbfId"`
	Name        string   `json:"name"`
	Text        string   `json:"text"`
	Flavor      string   `json:"flavor"`
	Artist      string   `json:"artist"`
	Attack      int      `json:"attack"`
	CardClass   string   `json:"cardClass"`
	Collectible bool     `json:"collectible"`
	Cost        int      `json:"cost"`
	Elite       bool     `json:"elite"`
	Faction     string   `json:"faction"`
	Health      int      `json:"health"`
	Mechanics   []string `json:"mechanics"`
	Rarity      string   `json:"rarity"`
	Set         string   `json:"set"`
	Type        string   `json:"type"`
}
