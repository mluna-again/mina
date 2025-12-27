package main

import (
	"cmp"
	"fmt"
	"slices"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
)

func (m model) fzfView() string {
	list := m.theme.list.Render(m.list.View())
	list = lipgloss.Place(m.list.Width(), m.list.Height(), lipgloss.Center, lipgloss.Center, list, lipgloss.WithWhitespaceBackground(m.theme.list.GetBackground()))

	return m.promptView() + "\n" + list
}

func (m model) promptView() string {
	bg := m.tinput.TextStyle.GetBackground()
	title := m.theme.title.Render(fmt.Sprintf(" %s ", m.title))
	title = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, title, lipgloss.WithWhitespaceBackground(bg))

	input := m.tinput.View()
	input = lipgloss.PlaceHorizontal(m.width, lipgloss.Center, input, lipgloss.WithWhitespaceBackground(bg))

	fill := lipgloss.PlaceHorizontal(m.width, lipgloss.Center, "", lipgloss.WithWhitespaceBackground(bg))

	return lipgloss.JoinVertical(lipgloss.Top, title, input, fill)
}

func (m model) confirmView() string {
	// user set it
	h := 3
	if m.ignoreHeight {
		h = m.height
	}
	w := m.width
	if m.ignoreWidth {
		w = m.width
	}

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

	bg := m.tinput.PromptStyle.GetBackground()
	input = lipgloss.Place(int(float64(m.width)*0.2), h, lipgloss.Center, lipgloss.Center, input, lipgloss.WithWhitespaceBackground(bg))

	msg := lipgloss.Place(w-lipgloss.Width(input), h, lipgloss.Center, lipgloss.Center, m.title)
	msg = m.theme.title.Render(msg)

	content := lipgloss.JoinHorizontal(lipgloss.Left, msg, input)
	return content
}

func (m model) menuView() string {
	bg := m.theme.bg.GetBackground()

	fill := lipgloss.PlaceHorizontal(m.width, lipgloss.Center, "", lipgloss.WithWhitespaceBackground(bg))

	titleStr := m.theme.title.Render(fmt.Sprintf(" %s ", m.title))
	title := lipgloss.PlaceHorizontal(m.width, lipgloss.Center, titleStr, lipgloss.WithWhitespaceBackground(bg))

	items := []MappedKey{}
	for line, key := range m.menuKeys {
		keyStr := m.theme.bg.Render(fmt.Sprintf(" %s", line))
		actionStr := m.theme.menuKey.Render(fmt.Sprintf("[%s] ", key.action))
		keyPadded := lipgloss.PlaceHorizontal(m.width-lipgloss.Width(actionStr), lipgloss.Left, keyStr, lipgloss.WithWhitespaceBackground(bg))

		item := lipgloss.JoinHorizontal(lipgloss.Left, keyPadded, actionStr)
		items = append(items, MappedKey{str: item, index: key.index})
	}
	slices.SortFunc[[]MappedKey](items, func(prev MappedKey, next MappedKey) int {
		return cmp.Compare[int](prev.index, next.index)
	})

	itemsStr := []string{}
	for _, item := range items {
		itemsStr = append(itemsStr, item.str)
	}

	if len(items) == 0 {
		return title
	}

	itemsJoined := lipgloss.JoinVertical(lipgloss.Top, itemsStr...)
	content := lipgloss.JoinVertical(lipgloss.Top, title, fill, itemsJoined, fill)

	return content
}
