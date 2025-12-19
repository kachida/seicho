import (
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var desc, freq string
var tags []string

var alarmCommand = &cobra.Command{
	Use:   "alarm [title]",
	Short: "Add a new alarm",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data := loadData()
		newAlarm := Alarm{
			ID:          len(data.Alarms) + 1,
			Title:       args[0],
			Description: desc,
			CreatedAt:   time.Now().Format("2006-01-02"),
			Frequency:   freq,
			Tags:        tags,
			Status:      "active",
		}
		data.Alarms = append(data.Alarms, newAlarm)
		saveData(data)
		color.Green("✅ Added alarm: %s", newAlarm.Title)
	},
}

func init() {
	alarmCommand.Flags().StringVarP(&desc, "desc", "d", "", "Description of the alarm")
	alarmCommand.Flags().StringVarP(&freq, "freq", "f", "once", "Frequency (once/daily/weekly)")
	alarmCommand.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, "Comma-separated tags")
	rootCmd.AddCommand(alarmCommand)
}
