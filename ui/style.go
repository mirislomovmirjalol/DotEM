package ui

import "github.com/charmbracelet/lipgloss"

var unfocused = lipgloss.NewStyle()
var focused = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#a78bfa"))
var navigationHelper = "" +
	"Up / Down: Navigate | Enter: Select | Backspace: Delete | Ctrl+C: Quit"

var title = lipgloss.NewStyle().Padding(1, 8).Background(lipgloss.Color("#7e22ce"))
