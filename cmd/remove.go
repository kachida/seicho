package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove [goalID]",
	Short: "Remove a goal by its ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var id int
		fmt.Sscanf(args[0], "%d", &id)
		data := loadData()
		newGoals := []Goal{}
		found := false
		for _, g := range data.Goals {
			if g.ID != id {
				newGoals = append(newGoals, g)
			} else {
				found = true
			}
		}
		if !found {
			color.Red("Goal #%d not found.", id)
			return
		}
		data.Goals = newGoals
		saveData(data)
		color.Green("🗑️  Removed goal #%d", id)
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}