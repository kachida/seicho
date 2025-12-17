package cmd

import "github.com/spf13/cobra"

var motivateCmd = &cobra.Command{
	Use:   "motivate",
	Short: "Display a random inspirational quote",
	Run: func(cmd *cobra.Command, args []string) {
		showQuote()
	},
}

func init() {
	rootCmd.AddCommand(motivateCmd)
}