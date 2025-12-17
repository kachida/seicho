package cmd

import (
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var desc, freq string
var tags []string


var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a new goal",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data := loadData()
		newGoal := Goal{
			ID:              len(data.Goals) + 1,
			Title:           args[0],
			Description:     desc,
			CreatedAt:       time.Now().Format("2006-01-02"),
			TargetFrequency: freq,
			Tags:            tags,
			Progress:        []Progress{},
		}
		data.Goals = append(data.Goals, newGoal)
		saveData(data)
		color.Green("✅ Added goal: %s", newGoal.Title)
	},
	
}

func init() {
	addCmd.Flags().StringVarP(&desc, "desc", "d", "", "Description of the goal")
	addCmd.Flags().StringVarP(&freq, "freq", "f", "daily", "Target frequency (daily/weekly)")
	addCmd.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, "Comma‑separated tags")
	rootCmd.AddCommand(addCmd)
}