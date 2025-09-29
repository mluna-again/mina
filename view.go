package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
)

func (m model) headerView() string {
	bg := m.tinput.TextStyle.GetBackground()
	title := m.theme.title.Render(fmt.Sprintf(" %s ", m.title))
	title = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, title, lipgloss.WithWhitespaceBackground(bg))

	input := m.tinput.View()
	input = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, input, lipgloss.WithWhitespaceBackground(bg))

	fill := lipgloss.PlaceHorizontal(m.width, lipgloss.Center, "", lipgloss.WithWhitespaceBackground(bg))

	return fmt.Sprintf("%s\n%s\n%s", title, input, fill)
}

func (m model) listView() string {
	list := m.theme.list.Render(m.list.View())
	list = lipgloss.Place(m.list.Width(), m.list.Height(), lipgloss.Center, lipgloss.Center, list, lipgloss.WithWhitespaceBackground(m.theme.list.GetBackground()))

	return list
}

func (m model) confirmView() string {
	input := m.tinput.Value()
	if input == "" {
		input = m.theme.placeholder.Render("N/y")
	} else if (utf8.RuneCount([]byte(input))) == 1 {
		input = m.theme.prompt.Render(fmt.Sprintf(" %s ", input))
	} else if (utf8.RuneCount([]byte(input))) == 2 {
		input = m.theme.prompt.Render(fmt.Sprintf("%s ", input))
	} else {
		input = m.theme.prompt.Render(input)
	}

	input = lipgloss.PlaceHorizontal(int(float64(m.width)*0.2), lipgloss.Center, input, lipgloss.WithWhitespaceBackground(m.tinput.PromptStyle.GetBackground()))

	msg := lipgloss.PlaceHorizontal(m.width-lipgloss.Width(input), lipgloss.Center, m.title)
	msg = m.theme.title.Render(msg)

	return lipgloss.JoinHorizontal(lipgloss.Left, msg, input)
}
