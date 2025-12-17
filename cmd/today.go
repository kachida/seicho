package cmd

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Show today’s progress summary",
	Run: func(cmd *cobra.Command, args []string) {
		data := loadData()
		today := time.Now().Format("2006-01-02")
		color.Green("\nToday's Progress (%s):\n", today)
		for _, g := range data.Goals {
			for _, p := range g.Progress {
				if p.Date == today {
					fmt.Printf("[%d] %s — %s\n", g.ID, g.Title, p.Status)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}