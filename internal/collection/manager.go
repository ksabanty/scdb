package collection

import (
	"errors"
	"scdb/internal/models"
	"scdb/internal/storage"
)

func AddCard(card models.Card) error {
	cards, err := storage.GetCards()
	if err != nil {
		return err
	}
	cards = append(cards, card)
	return storage.SaveCards(cards)
}

func DeleteCard(player string) error {
	cards, err := storage.GetCards()
	if err != nil {
		return err
	}

	newCards := []models.Card{}
	found := false
	for _, c := range cards {
		if c.Player != player {
			newCards = append(newCards, c)
		} else {
			found = true
		}
	}

	if !found {
		return errors.New("card not found")
	}

	return storage.SaveCards(newCards)
}

func ListCards() ([]models.Card, error) {
	return storage.GetCards()
}
