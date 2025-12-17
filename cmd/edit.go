package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var newTitle, newDesc string

var editCmd = &cobra.Command{
	Use:   "edit [goalID]",
	Short: "Edit a goal’s title or description",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var id int
		fmt.Sscanf(args[0], "%d", &id)
		data := loadData()
		for i, g := range data.Goals {
			if g.ID == id {
				if newTitle != "" {
					data.Goals[i].Title = newTitle
				}
				if newDesc != "" {
					data.Goals[i].Description = newDesc
				}
				saveData(data)
				color.Green("✏️  Updated goal #%d", id)
				return
			}
		}
		color.Red("Goal not found.")
	},
}

func init() {
	editCmd.Flags().StringVarP(&newTitle, "title", "t", "", "New title")
	editCmd.Flags().StringVarP(&newDesc, "desc", "d", "", "New description")
	rootCmd.AddCommand(editCmd)
}