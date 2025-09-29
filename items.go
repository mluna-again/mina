package main

import (
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type item struct {
	text          string
	style         lipgloss.Style
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

	line := lipgloss.PlaceHorizontal(m.Width(), lipgloss.Left, getColumns(i.text, separator, displayColumns))
	if index == m.Index() {
		line = i.selectedStyle.Render(line)
	} else {
		line = i.style.Render(line)
	}

	fmt.Fprint(w, line)
}

func getColumns(line string, sep string, cols string) string {
	if cols == "" {
		return line
	}

	start := -1
	end := -1
	fixed := -1

	var err error
	if strings.Contains(cols, ",") {
		cmps := strings.Split(cols, ",")
		start, err = strconv.Atoi(cmps[0])
		if err != nil {
			return line
		}

		end, err = strconv.Atoi(cmps[1])
		if err != nil {
			return line
		}
	} else if cols != "" {
		fixed, err = strconv.Atoi(cols)
		if err != nil {
			return line
		}
	}

	filtered := []string{}
	for i, cmp := range strings.Split(line, sep) {
		if fixed != -1 && i != fixed {
			continue
		}

		if start != -1 && i < start {
			continue
		}

		if end != -1 && i > end {
			continue
		}

		filtered = append(filtered, cmp)
	}

	return strings.Join(filtered, sep)
}
