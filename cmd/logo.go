package cmd

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var (
	green   = color.New(color.FgHiGreen).SprintFunc()
	cyan    = color.New(color.FgCyan).SprintFunc()
	white   = color.New(color.FgWhite).SprintFunc()
	magenta = color.New(color.FgMagenta).SprintFunc()
)

func ShowAnimatedLogo() {
	logoLines := []string{
" __      _      _           ",
"/ _\\ ___(_) ___| |__   ___  ",
"\\ \\ / _ \\ |/ __| '_ \\ / _ \\ ",
"_\\ \\  __/ | (__| | | | (_) |",
"\\__/\\___|_|\\___|_| |_|\\___/ ",
"                            ",                     
	}

	tagline := "Seicho — Growth in Action"
	quote := "\"Small steps, taken daily, lead to great growth.\""

	for i, line := range logoLines {
		switch {
		case i < 2:
			fmt.Println(green(line))
		case i < 4:
			fmt.Println(cyan(line))
		default:
			fmt.Println(white(line))
		}
		time.Sleep(120 * time.Millisecond)
	}

	fmt.Println(cyan("\n   " + tagline))
	fmt.Println(magenta("   " + quote + "\n"))
}