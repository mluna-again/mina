package main

import (
	"github.com/charmbracelet/lipgloss"
)

type theme struct {
	// TODO: fix these names (add fg/bg variants)
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

func loadTheme(t ConfigTheme) theme {
	theme := kanagawaDragon()

	if t.TitleBG != "" {
		theme.title = theme.title.Background(lipgloss.Color(t.TitleBG))
	}

	if t.TitleFG != "" {
		theme.title = theme.title.Foreground(lipgloss.Color(t.TitleFG))
	}

	if t.ListItemBG != "" {
		theme.listItem = theme.listItem.Background(lipgloss.Color(t.ListItemBG))
	}

	if t.ListItemFG != "" {
		theme.listItem = theme.listItem.Foreground(lipgloss.Color(t.ListItemFG))
	}

	if t.SelectedListItemBG != "" {
		theme.selectedListItem = theme.selectedListItem.Background(lipgloss.Color(t.SelectedListItemBG))
	}

	if t.SelectedListItemFG != "" {
		theme.selectedListItem = theme.selectedListItem.Foreground(lipgloss.Color(t.SelectedListItemFG))
	}

	return theme
}

func kanagawaDragon() theme {
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
