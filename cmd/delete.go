package cmd

import (
	"fmt"
	"scdb/internal/collection"
	"scdb/internal/models"

	"github.com/spf13/cobra"
)

var (
	deletePlayer       string
	deleteTeam         string
	deleteManufacturer string
	deleteCollection   string
	deleteYear         string
	deleteValue        string
	deleteRookie       bool
	deleteNumbered     bool
	deleteAuto         bool
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a card from your collection by player name",
	Run: func(cmd *cobra.Command, args []string) {
		filter := models.Card{
			Player:       deletePlayer,
			Team:         deleteTeam,
			Manufacturer: deleteManufacturer,
			Collection:   deleteCollection,
			Year:         deleteYear,
			Value:        deleteValue,
			Rookie:       deleteRookie,
			Numbered:     deleteNumbered,
			Auto:         deleteAuto,
		}
		if err := collection.DeleteCard(filter); err != nil {
			fmt.Println("Error deleting card:", err)
			return
		}
		fmt.Println("Card deleted successfully!")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVar(&deletePlayer, "player", "", "Player name of the card to delete")
	deleteCmd.Flags().StringVar(&deleteTeam, "team", "", "Team name")
	deleteCmd.Flags().StringVar(&deleteManufacturer, "manufacturer", "", "Card manufacturer")
	deleteCmd.Flags().StringVar(&deleteCollection, "collection", "", "Collection name")
	deleteCmd.Flags().StringVar(&deleteYear, "year", "", "Year of the card")
	deleteCmd.Flags().StringVar(&deleteValue, "value", "", "Card value")
	deleteCmd.Flags().BoolVar(&deleteRookie, "rookie", false, "Is rookie card")
	deleteCmd.Flags().BoolVar(&deleteNumbered, "numbered", false, "Is numbered card")
	deleteCmd.Flags().BoolVar(&deleteAuto, "auto", false, "Is autographed")
	deleteCmd.MarkFlagRequired("player")
}
