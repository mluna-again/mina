package main

import "github.com/charmbracelet/lipgloss"

type theme struct {
	title            lipgloss.Style
	bg               lipgloss.Style
	menuKey          lipgloss.Style
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
	hidden := lipgloss.NewStyle().Background(lipgloss.Color("0")).Foreground(lipgloss.Color("0"))

	p := lipgloss.NewStyle().
		Background(lipgloss.Color("0")).
		Foreground(lipgloss.Color("15"))

	ph := lipgloss.NewStyle().
		Background(lipgloss.Color("0")).
		Foreground(lipgloss.Color("5"))

	pC := lipgloss.NewStyle().
		Foreground(lipgloss.Color("15")).
		Background(lipgloss.Color("15"))

	t := lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")).
		Background(lipgloss.Color("3"))

	l := lipgloss.NewStyle().
		Foreground(lipgloss.Color("15")).
		Background(lipgloss.Color("233")).
		Padding(0, 0)

	li := lipgloss.NewStyle().
		Foreground(lipgloss.Color("15")).
		Background(lipgloss.Color("233")).
		Padding(0, 1)

	sli := lipgloss.NewStyle().
		Foreground(lipgloss.Color("0")).
		Background(lipgloss.Color("3")).
		Padding(0, 1)

	mk := lipgloss.NewStyle().
		Background(lipgloss.Color("0")).
		Foreground(lipgloss.Color("1"))

	return theme{
		bg:               p,
		menuKey:          mk,
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
