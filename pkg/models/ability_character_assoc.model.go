package models

type AbilityCharacterAssoc struct {
    Id          string `json:"ability_id"`
    CharacterID string `json:"character_id"`
    Value       int    `json:"value"`
}
