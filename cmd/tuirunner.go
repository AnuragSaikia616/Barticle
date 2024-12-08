package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func RunBarticleTUI() {
	m := InitMainModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Errorf("Error running the TUI\nError: %v", err)
		os.Exit(1)
	}
}
