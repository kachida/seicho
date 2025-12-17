package cmd

import (
	"encoding/json"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)


var dataFile = filepath.Join(os.Getenv("HOME"), ".seicho", "data.json")
var quotesFile = filepath.Join(os.Getenv("HOME"), ".seicho", "quotes.json")

var rootCmd = &cobra.Command{
	Use:   "seicho",
	Short: "Seicho — Track your growth and consistency with inspiration",
	Long:  "Seicho helps you measure progress, stay consistent, and grow daily.",
	Run: func(cmd *cobra.Command, args []string) {
		ShowAnimatedLogo()
		showQuote()
		listGoals()
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

type Progress struct {
	Date   string `json:"date"`
	Status string `json:"status"`
}

type Goal struct {
	ID              int        `json:"id"`
	Title           string     `json:"title"`
	Description     string     `json:"description"`
	CreatedAt       string     `json:"created_at"`
	TargetFrequency string     `json:"target_frequency"`
	Tags            []string   `json:"tags"`
	Progress        []Progress `json:"progress"`
}

type Data struct {
	Goals []Goal `json:"goals"`
}

func loadData() Data {
	var d Data
	b, err := os.ReadFile(dataFile)
	if err != nil {
		return Data{}
	}
	json.Unmarshal(b, &d)
	return d
}

func saveData(d Data) {
	os.MkdirAll(filepath.Dir(dataFile), 0755)
	b, _ := json.MarshalIndent(d, "", "  ")
	os.WriteFile(dataFile, b, 0644)
}

func showQuote() {
	b, err := os.ReadFile(quotesFile)
	if err != nil {
		return
	}
	var q struct{ Quotes []string `json:"quotes"` }
	json.Unmarshal(b, &q)
	if len(q.Quotes) == 0 {
		return
	}
	rand.Seed(time.Now().UnixNano())
	color.Cyan("\n💡 %s\n", q.Quotes[rand.Intn(len(q.Quotes))])
}

