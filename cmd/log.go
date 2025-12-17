package cmd

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log [goalID] [done|missed]",
	Short: "Log progress for a goal",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		data := loadData()
		var id int
		fmt.Sscanf(args[0], "%d", &id)
		status := args[1]

		for i, g := range data.Goals {
			if g.ID == id {
				data.Goals[i].Progress = append(g.Progress, Progress{
					Date:   time.Now().Format("2006-01-02"),
					Status: status,
				})
				saveData(data)
				color.Green("📅 Logged %s for goal #%d", status, id)
				return
			}
		}
		color.Red("Goal not found.")
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}