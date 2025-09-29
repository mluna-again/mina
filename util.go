package main

import (
	"io"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) calculatePromptWidth() int {
	return m.width - lipgloss.Width(m.tinput.Prompt) - 3 // padding + cursor
}

func loadInput() ([]string, error) {
	lines, err := io.ReadAll(os.Stdin)
	if err != nil {
		return []string{}, err
	}

	filtered := []string{}
	for _, line := range strings.Split(string(lines), "\n") {
		if line == "" {
			continue
		}

		filtered = append(filtered, line)
	}

	return filtered, nil
}

func isFilterEvent(msg tea.Msg) bool {
	t, ok := msg.(tea.KeyMsg)
	if !ok {
		return false
	}

	key := t.String()

	return key != "ctrl+n" && key != "ctrl+p" && key != "enter"
}
