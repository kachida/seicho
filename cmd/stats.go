package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show success and consistency metrics",
	Run: func(cmd *cobra.Command, args []string) {
		data := loadData()
		if len(data.Goals) == 0 {
			color.Yellow("No goals to show stats for.")
			return
		}
		color.Green("\nGoal Stats:\n")
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
			fmt.Printf("[%d] %s — %d%% success (%d/%d days)\n", g.ID, g.Title, success, done, total)
		}
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}