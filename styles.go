package main

import "github.com/charmbracelet/lipgloss"

type theme struct {
	title            lipgloss.Style
	bg               lipgloss.Style
	prompt           lipgloss.Style
	promptCursor     lipgloss.Style
	placeholder      lipgloss.Style
	noCursor         lipgloss.Style
	list             lipgloss.Style
	listItem         lipgloss.Style
	selectedListItem lipgloss.Style
}

// Kanagawa Dragon
func getTheme() theme {
	hidden := lipgloss.NewStyle().Background(lipgloss.Color("#282727")).Foreground(lipgloss.Color("#282727"))

	p := lipgloss.NewStyle().
		Background(lipgloss.Color("#282727")).
		Foreground(lipgloss.Color("#c5c9c5"))

	ph := lipgloss.NewStyle().
		Background(lipgloss.Color("#282727")).
		Foreground(lipgloss.Color("#a6a69c"))

	pC := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#c5c9c5")).
		Background(lipgloss.Color("#c5c9c5"))

	t := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#181616")).
		Background(lipgloss.Color("#c4746e"))

	l := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#c5c9c5")).
		Background(lipgloss.Color("#12120f")).
		Padding(0, 0)

	li := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#c5c9c5")).
		Background(lipgloss.Color("#12120f")).
		Padding(0, 1)

	sli := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#c5c9c5")).
		Background(lipgloss.Color("#c4746e")).
		Padding(0, 1)

	return theme{
		bg:               p,
		prompt:           p,
		promptCursor:     pC,
		noCursor:         hidden,
		placeholder:      ph,
		title:            t,
		list:             l,
		listItem:         li,
		selectedListItem: sli,
	}
}
