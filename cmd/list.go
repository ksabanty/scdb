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
	Run:   listRun,
}

func listRun(cmd *cobra.Command, args []string) {
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

	var totalValue float64
	for _, c := range cards {
		valueFormatted := c.Value
		if c.Value != "" {
			var amount float64
			_, err := fmt.Sscanf(c.Value, "%f", &amount)
			if err == nil {
				valueFormatted = fmt.Sprintf("$%.2f", amount)
				totalValue += amount
			}
		}
		row := []string{
			c.Player,
			c.Team,
			c.Manufacturer,
			c.Collection,
			c.Year,
			valueFormatted,
			boolToYN(c.Rookie),
			boolToYN(c.Numbered),
			boolToYN(c.Auto),
		}
		table.Append(row)
	}

	// Add an empty row
	table.Append([]string{"", "", "", "", "", "", "", "", ""})

	// Add total value row, aligned with the Value column (column 6)
	totalRow := []string{"", "", "", "", "", fmt.Sprintf("$%.2f", totalValue), "", "", ""}
	table.Append(totalRow)

	table.Render()
}

// boolToYN returns "Y" for true and "N" for false
func boolToYN(b bool) string {
	if b {
		return "Y"
	}
	return "N"
}

func init() {
	rootCmd.AddCommand(listCmd)
}
