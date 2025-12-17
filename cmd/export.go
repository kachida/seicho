package cmd

import (
	"encoding/csv"
	"os"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export [filename]",
	Short: "Export all goal data to a CSV file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data := loadData()
		file, err := os.Create(args[0])
		if err != nil {
			color.Red("Error creating file: %v", err)
			return
		}
		defer file.Close()
		writer := csv.NewWriter(file)
		defer writer.Flush()

		writer.Write([]string{"ID", "Title", "Date", "Status"})
		for _, g := range data.Goals {
			for _, p := range g.Progress {
				writer.Write([]string{
					fmt.Sprintf("%d", g.ID),
					g.Title,
					p.Date,
					p.Status,
				})
			}
		}
		color.Green("📤 Exported data to %s", args[0])
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}