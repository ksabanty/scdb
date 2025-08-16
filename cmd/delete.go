package cmd

import (
	"fmt"
	"scdb/internal/collection"

	"github.com/spf13/cobra"
)

var deletePlayer string

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a card from your collection by player name",
	Run: func(cmd *cobra.Command, args []string) {
		if err := collection.DeleteCard(deletePlayer); err != nil {
			fmt.Println("Error deleting card:", err)
			return
		}
		fmt.Printf("Card(s) for player '%s' deleted successfully!\n", deletePlayer)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVar(&deletePlayer, "player", "", "Player name of the card to delete")
	deleteCmd.MarkFlagRequired("player")
}
