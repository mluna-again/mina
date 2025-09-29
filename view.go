package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m model) header() string {
	bg := m.tinput.TextStyle.GetBackground()
	title := m.theme.title.Render(fmt.Sprintf(" %s ", m.title))
	title = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, title, lipgloss.WithWhitespaceBackground(bg))

	input := m.tinput.View()
	input = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, input, lipgloss.WithWhitespaceBackground(bg))

	fill := lipgloss.PlaceHorizontal(m.width, lipgloss.Center, "", lipgloss.WithWhitespaceBackground(bg))

	return fmt.Sprintf("%s\n%s\n%s\n", title, input, fill)
}
