package main

import "github.com/charmbracelet/lipgloss"

type theme struct {
	title            lipgloss.Style
	prompt           lipgloss.Style
	promptCursor     lipgloss.Style
	listItem         lipgloss.Style
	selectedListItem lipgloss.Style
}

// Kanagawa Dragon
func getTheme() theme {
	p := lipgloss.NewStyle().
		Background(lipgloss.Color("#282727")).
		Foreground(lipgloss.Color("#c5c9c5"))
	pC := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#c5c9c5")).
		Background(lipgloss.Color("#c5c9c5"))

	t := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#181616")).
		Background(lipgloss.Color("#c4746e"))

	li := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#c5c9c5")).
		Background(lipgloss.Color("#1d1c19"))

	sli := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#c5c9c5")).
		Background(lipgloss.Color("#c4746e"))

	return theme{
		prompt:           p,
		promptCursor:     pC,
		title:            t,
		listItem:         li,
		selectedListItem: sli,
	}
}
