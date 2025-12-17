package cmd

import "github.com/spf13/cobra"

var zenCmd = &cobra.Command{
	Use:   "zen",
	Short: "Display the Seicho logo and a random inspirational quote",
	Run: func(cmd *cobra.Command, args []string) {
		ShowAnimatedLogo()
		showQuote()
	},
}

func init() {
	rootCmd.AddCommand(zenCmd)
}