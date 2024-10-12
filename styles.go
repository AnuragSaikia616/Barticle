package main

import "github.com/charmbracelet/lipgloss"

const (
	pink  = "#ffaadd"
	blue  = "#22ddff"
	green = "#33ff99"
)

var (
	serverStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(pink))
	questionStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(blue))
	answerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(green))
)
