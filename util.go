package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) calculatePromptWidth() int {
	if m.mode == CONFIRM_MODE {
		return 3
	}

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

func getMenuItems(input []string) map[string]Key {
	keys := map[string]Key{}
	j := 1

	for i, line := range input {
		cmps := strings.Split(line, separator)
		if len(cmps) < 0 || len(cmps) > 2 {
			log.Fatalf("%s: bad argument, separator: %s\n", line, separator)
		}

		if len(cmps) == 1 {
			keys[line] = Key{
				action: fmt.Sprintf("%d", j),
				index:  i,
				text:   line,
			}
			j++
			continue
		}

		keys[cmps[0]] = Key{
			action: cmps[1],
			index:  i,
			text:   cmps[0],
		}
	}

	return keys
}
