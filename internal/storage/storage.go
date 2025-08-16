package storage

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"scdb/internal/models"
)

const filePath = "data/collection.json"

func loadCards() ([]models.Card, error) {
	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return []models.Card{}, nil
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var cards []models.Card
	err = json.Unmarshal(data, &cards)
	return cards, err
}

func saveCards(cards []models.Card) error {
	data, err := json.MarshalIndent(cards, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

func GetCards() ([]models.Card, error) {
	return loadCards()
}

func SaveCards(cards []models.Card) error {
	return saveCards(cards)
}
