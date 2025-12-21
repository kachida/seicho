import (
	"time"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var goalTitle string

var streakCommand = &cobra.Command{
	Use:   "streak [goal]",
	Short: "Log progress for a goal and update your streak",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data := loadData()
		goalTitle := args[0]
		found := false

		for i, g := range data.Goals {
			if g.Title == goalTitle {
				found = true
				today := time.Now().Format("2006-01-02")

				// If last logged date is yesterday, increment streak
				lastDate, _ := time.Parse("2006-01-02", g.LastLogged)
				if g.LastLogged == today {
					color.Yellow("⚠️ Already logged progress for today!")
					return
				} else if lastDate.Add(24 * time.Hour).Format("2006-01-02") == today {
					g.Streak++
				} else {
					g.Streak = 1 // reset streak
				}

				g.LastLogged = today
				data.Goals[i] = g
				saveData(data)

				color.Green("🔥 Logged progress for '%s'!", g.Title)
				color.Cyan("Current streak: %d days", g.Streak)
				return
			}
		}

		if !found {
			color.Red("Goal not found: %s", goalTitle)
		}
	},
}

func init() {
	rootCmd.AddCommand(streakCommand)
}