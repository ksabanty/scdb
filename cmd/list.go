package cmd

import (
	"fmt"
	"os"
	"scdb/internal/collection"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all cards in your collection",
	Run: func(cmd *cobra.Command, args []string) {
		cards, err := collection.ListCards()
		if err != nil {
			fmt.Println("Error listing cards:", err)
			return
		}

		if len(cards) == 0 {
			fmt.Println("Your collection is empty.")
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.Header([]string{
			"Player", "Team", "Manufacturer", "Collection", "Year", "Value",
			"Rookie", "Numbered", "Auto",
		})

		for _, c := range cards {
			row := []string{
				c.Player,
				c.Team,
				c.Manufacturer,
				c.Collection,
				c.Year,
				c.Value,
				fmt.Sprintf("%t", c.Rookie),
				fmt.Sprintf("%t", c.Numbered),
				fmt.Sprintf("%t", c.Auto),
			}
			table.Append(row)
		}

		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
