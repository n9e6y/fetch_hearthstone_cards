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

type CollectableCard struct {
	Card
	Classes                        []string `json:"classes"`
	Armor                          int      `json:"armor"`
	HasDiamondSkin                 bool     `json:"hasDiamondSkin"`
	MercenariesRole                int      `json:"mercenariesRole"`
	HideCost                       bool     `json:"hideCost"`
	BattlegroundsBuddyDbfId        int      `json:"battlegroundsBuddyDbfId"`
	BattlegroundsDarkmoonPrizeTurn int      `json:"battlegroundsDarkmoonPrizeTurn"`
	Race                           string   `json:"race"`
	IsBattlegroundsPoolSpell       bool     `json:"isBattlegroundsPoolSpell"`
	TechLevel                      int      `json:"techLevel"`
	HideStats                      bool     `json:"hideStats"`
	TargetingArrowText             string   `json:"targetingArrowText"`
	BattlegroundsPremiumDbfId      int      `json:"battlegroundsPremiumDbfId"`
	CollectionText                 string   `json:"collectionText"`
	ReferencedTags                 []string `json:"referencedTags"`
	BattlegroundsNormalDbfId       int      `json:"battlegroundsNormalDbfId"`
	BattlegroundsSkinParentId      int      `json:"battlegroundsSkinParentId"`
	PuzzleType                     int      `json:"puzzleType"`
	Races                          []string `json:"races"`
	IsMiniSet                      bool     `json:"isMiniSet"`
	BattlegroundsHero              bool     `json:"battlegroundsHero"`
	MultiClassGroup                string   `json:"multiClassGroup"`
	HowToEarnGolden                string   `json:"howToEarnGolden"`
	SpellDamage                    int      `json:"spellDamage"`
	IsBattlegroundsBuddy           bool     `json:"isBattlegroundsBuddy"`
	Durability                     int      `json:"durability"`
	IsBattlegroundsPoolMinion      bool     `json:"isBattlegroundsPoolMinion"`
	MercenariesAbilityCooldown     int      `json:"mercenariesAbilityCooldown"`
	HowToEarn                      string   `json:"howToEarn"`
	CountAsCopyOfDbfId             int      `json:"countAsCopyOfDbfId"`
	SpellSchool                    string   `json:"spellSchool"`
	QuestReward                    string   `json:"questReward"`
	HeroPowerDbfId                 int      `json:"heroPowerDbfId"`
	Overload                       int      `json:"overload"`
}
