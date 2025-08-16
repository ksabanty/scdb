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

// DeleteCard deletes a card matching all non-empty fields in the filter.
func DeleteCard(filter models.Card) error {
	cards, err := storage.GetCards()
	if err != nil {
		return err
	}

	newCards := []models.Card{}
	found := false
	for _, c := range cards {
		if !cardMatchesFilter(c, filter) {
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

// cardMatchesFilter returns true if all non-zero fields in filter match the card
func cardMatchesFilter(card, filter models.Card) bool {
	if filter.Player != "" && card.Player != filter.Player {
		return false
	}
	if filter.Team != "" && card.Team != filter.Team {
		return false
	}
	if filter.Manufacturer != "" && card.Manufacturer != filter.Manufacturer {
		return false
	}
	if filter.Collection != "" && card.Collection != filter.Collection {
		return false
	}
	if filter.Year != "" && card.Year != filter.Year {
		return false
	}
	if filter.Value != "" && card.Value != filter.Value {
		return false
	}
	if filter.Rookie && !card.Rookie {
		return false
	}
	if filter.Numbered && !card.Numbered {
		return false
	}
	if filter.Auto && !card.Auto {
		return false
	}
	return true
}

func ListCards() ([]models.Card, error) {
	return storage.GetCards()
}
