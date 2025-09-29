package main

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item struct {
	text  string
	style lipgloss.Style
	selectedStyle lipgloss.Style
}

func (i item) FilterValue() string { return i.text }

type itemDelegate struct{}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	line := lipgloss.PlaceHorizontal(m.Width(), lipgloss.Left, i.text)
	if index == m.Index() {
		line = i.selectedStyle.Render(line)
	} else {
		line = i.style.Render(line)
	}

	fmt.Fprint(w, line)
}
