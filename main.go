package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

func main() {
	if os.Getenv("DEBUG") != "" {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Could not open log file: %v\n", err)
		} else {
			defer f.Close()
		}
	}

	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
