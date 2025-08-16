package cmd

import (
	"fmt"
	"scdb/internal/collection"
	"scdb/internal/models"

	"github.com/spf13/cobra"
)

var (
	player string
	team   string
	manuf  string
	coll   string
	yr     string
	val    string
	r      bool
	num    bool
	auto   bool
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new card to your collection",
	Run: func(cmd *cobra.Command, args []string) {
		if player == "" || team == "" {
			fmt.Println("Error: Both player and team fields are required.")
			return
		}
		card := models.Card{
			Player:       player,
			Team:         team,
			Manufacturer: manuf,
			Collection:   coll,
			Year:         yr,
			Value:        val,
			Rookie:       r,
			Numbered:     num,
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
	addCmd.Flags().StringVar(&manuf, "manufacturer", "", "Card manufacturer")
	addCmd.Flags().StringVar(&coll, "collection", "", "Collection name")
	addCmd.Flags().StringVar(&yr, "year", "", "Year")
	addCmd.Flags().StringVar(&val, "value", "", "value")
	addCmd.Flags().BoolVar(&r, "rookie", false, "Is rookie")
	addCmd.Flags().BoolVar(&num, "numbered", false, "Is numbered")
	addCmd.Flags().BoolVar(&auto, "auto", false, "Is autographed")

	addCmd.MarkFlagRequired("player")
	addCmd.MarkFlagRequired("team")
}
