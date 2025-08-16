package cmd

import (
	"fmt"
	"scdb/internal/collection"
	"scdb/internal/models"

	"github.com/spf13/cobra"
)

var (
	player         string
	team           string
	manufacturer   string
	collectionName string
	year           string
	value          string
	rookie         bool
	numbered       bool
	auto           bool
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new card to your collection",
	Run: func(cmd *cobra.Command, args []string) {
		card := models.Card{
			Player:       player,
			Team:         team,
			Manufacturer: manufacturer,
			Collection:   collectionName,
			Year:         year,
			Value:        value,
			Rookie:       rookie,
			Numbered:     numbered,
			Auto:         auto,
		}

		if err := collection.AddCard(card); err != nil {
			fmt.Println("Error adding card:", err)
			return
		}
		fmt.Println("Card added successfully!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVar(&player, "player", "", "Player name")
	addCmd.Flags().StringVar(&team, "team", "", "Team name")
	addCmd.Flags().StringVar(&manufacturer, "manufacturer", "", "Card manufacturer")
	addCmd.Flags().StringVar(&collectionName, "collection", "", "Collection name")
	addCmd.Flags().StringVar(&year, "year", "", "Year of the card")
	addCmd.Flags().StringVar(&value, "value", "", "Card value")
	addCmd.Flags().BoolVar(&rookie, "rookie", false, "Is rookie card")
	addCmd.Flags().BoolVar(&numbered, "numbered", false, "Is numbered card")
	addCmd.Flags().BoolVar(&auto, "auto", false, "Is autographed")

	addCmd.MarkFlagRequired("player")
	addCmd.MarkFlagRequired("team")
}
