package main

import tea "github.com/charmbracelet/bubbletea"

func (m model) updatePrompt(msg tea.Msg) (model, tea.Cmd) {
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

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, tea.Batch(cmds...)
}

func (m model) updateFzf(msg tea.Msg) (model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.tinput, cmd = m.tinput.Update(msg)
	cmds = append(cmds, cmd)

	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		if !m.ignoreHeight {
			m.height = msg.Height
			m.list.SetHeight(m.height - 3) // header
		}
		if !m.ignoreWidth {
			m.width = msg.Width
			m.list.SetWidth(m.width)
			m.tinput.Width = m.calculatePromptWidth()
		}
		return m, tea.ClearScreen

	case tea.KeyMsg:
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
			if len(m.content) > 1 {
				response = m.content[m.list.GlobalIndex()]
			}
			return m, tea.Quit

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

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

		case "ctrl+c", "q":
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
			if key.action == keyStroke {
				response = key.text
				return m, tea.Quit
			}
		}

		switch keyStroke {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}
