package cmd

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/lipgloss"
)

var InputStyle, _ = textarea.DefaultStyles()

var SummaryStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#ffdddd"))

var QuestionStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#ccbbff"))

var AnswerStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#ffffaa"))

var LogoStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#33ff99"))

const (
	logo = `>_< BARTicle...`
	desc = "\nA text summarization and question answering tool! (Press Ctrl+D to exit.)\n\n"
)
