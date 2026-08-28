package main

import (
	"os"
	"unicode"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) updatePrompt(msg tea.Msg) (model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.ignoreHeight {
			m.height = msg.Height
		}
		if !m.ignoreWidth {
			m.width = msg.Width
			m.tinput.Width = m.calculatePromptWidth()
		}
		return m, tea.ClearScreen

	case tea.KeyMsg:
		if key.Matches(msg, m.tinput.KeyMap.DeleteWordBackward) {
			m.deleteBackwards()
			return m, nil
		}

		switch msg.String() {
		case "enter":
			response = m.tinput.Value()
			return m, tea.Quit

		case "ctrl+c", "esc":
			return m, tea.Quit
		}
	}

	m.tinput, cmd = m.tinput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) updateFzf(msg tea.Msg) (model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.ignoreHeight {
			m.height = msg.Height
			h := 3
			if hideBanner {
				h = 1
			}
			m.list.SetHeight(m.height - h) // header
		}
		if !m.ignoreWidth {
			m.width = msg.Width
			m.list.SetWidth(m.width)
			m.tinput.Width = m.calculatePromptWidth()
		}
		return m, tea.ClearScreen

	case tea.KeyMsg:
		if key.Matches(msg, m.tinput.KeyMap.DeleteWordBackward) {
			m.deleteBackwards()
			return m, nil
		}
		// man this makes my brain hurt
		// why do i need this? i *THINK* its because if i set the filter on each keystroke
		// it resets the list position to 0. so if the keys are ctrl+n or ctrl+p i dont update the
		// filter to avoid it resetting the list.
		// also, i dont update the filter on enter because if i do that the index will reset before i set the final response
		if isFilterEvent(msg) {
			m.list.SetFilterText(m.tinput.Value())
		}
		switch msg.String() {
		case "enter":
			if len(m.list.VisibleItems()) > 0 {
				response = m.content[m.list.GlobalIndex()]
			}
			return m, tea.Quit

		case "ctrl+c", "esc":
			return m, tea.Quit
		}
	}

	m.tinput, cmd = m.tinput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) updateConfirm(msg tea.Msg) (model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.tinput, cmd = m.tinput.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.ignoreHeight {
			m.height = msg.Height
		}
		if !m.ignoreWidth {
			m.width = msg.Width
			m.tinput.Width = m.calculatePromptWidth()
		}
		return m, tea.ClearScreen

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			response = m.tinput.Value()
			return m, tea.Quit

		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		}

		if onekey {
			response = m.tinput.Value()
			return m, tea.Quit
		}
	}

	return m, tea.Batch(cmds...)
}

func (m model) updateMenu(msg tea.Msg) (model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.ignoreHeight {
			m.height = msg.Height
		}
		if !m.ignoreWidth {
			m.width = msg.Width
			m.tinput.Width = m.calculatePromptWidth()
		}
		return m, tea.ClearScreen

	case tea.KeyMsg:
		keyStroke := msg.String()
		for _, key := range m.menuKeys {
			// special case
			if key.action == "Space" && keyStroke == " " {
				response = key.text
				return m, tea.Quit
			}

			if key.action == keyStroke {
				response = key.text
				return m, tea.Quit
			}
		}

		switch keyStroke {
		case "ctrl+c", "esc":
			return m, tea.Quit
		}
	}

	return m, nil
}

// slightly modified input.DeleteWordBackwards but path aware
func (m *model) deleteBackwards() {
	if m.tinput.Position() == 0 || len(m.tinput.Value()) == 0 {
		return
	}

	if m.tinput.EchoMode != textinput.EchoNormal {
		m.tinput.SetValue("")
		return
	}

	oldPos := m.tinput.Position()

	m.tinput.SetCursor(m.tinput.Position() - 1)
	for unicode.IsSpace(rune(m.tinput.Value()[m.tinput.Position()])) {
		if m.tinput.Position() <= 0 {
			break
		}
		// ignore series of whitespace before cursor
		m.tinput.SetCursor(m.tinput.Position() - 1)
	}

	for m.tinput.Position() > 0 {
		c := m.tinput.Value()[m.tinput.Position()]
		if unicode.IsSpace(rune(c)) || c == os.PathSeparator {
			if m.tinput.Position() > 0 {
				// keep the previous space
				m.tinput.SetCursor(m.tinput.Position() + 1)
			}
			break
		}
		m.tinput.SetCursor(m.tinput.Position() - 1)
	}

	if oldPos > len(m.tinput.Value()) {
		m.tinput.SetValue(m.tinput.Value()[:m.tinput.Position()])
	} else {
		v := append([]rune(m.tinput.Value()[:m.tinput.Position()]), []rune(m.tinput.Value()[oldPos:])...)
		m.tinput.SetValue(string(v))
	}
}
