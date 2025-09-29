package main

import (
	"io"
	"os"
	"strings"

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

	return strings.Split(string(lines), "\n"), nil
}
