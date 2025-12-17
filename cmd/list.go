package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all goals",
	Run: func(cmd *cobra.Command, args []string) {
		showQuote()
		listGoals()
	},
}

func listGoals() {
	data := loadData()
	if len(data.Goals) == 0 {
		color.Yellow("No goals yet. Add one with: seicho add \"Title\" -d \"Description\" -f daily")
		return
	}
	color.Green("\nYour Goals:\n")
	for _, g := range data.Goals {
		done := 0
		for _, p := range g.Progress {
			if p.Status == "done" {
				done++
			}
		}
		total := len(g.Progress)
		success := 0
		if total > 0 {
			success = (done * 100) / total
		}
		fmt.Printf("[%d] %s — %d%% success\n", g.ID, g.Title, success)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
}