package main

import "github.com/charmbracelet/lipgloss"

func (m model) calculatePromptWidth() int {
	return m.width - lipgloss.Width(m.tinput.Prompt) - 3 // padding + cursor
}
