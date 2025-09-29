package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

var icon string
var title string

type model struct {
	theme  theme
	width  int
	height int
	tinput textinput.Model
	title  string
}

func newMina() model {
	theme := getTheme()

	t := textinput.New()
	t.Focus()
	t.Prompt = fmt.Sprintf("%s ", icon)
	t.Width = 80
	t.TextStyle = theme.prompt
	t.PromptStyle = theme.prompt
	// This disables the blinking that i *didnt* enable, btw
	t.Cursor.Style = theme.promptCursor
	t.Cursor.TextStyle = theme.promptCursor

	return model{
		tinput: t,
		title:  title,
		theme:  theme,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		m.tinput.Width = m.calculatePromptWidth()
		return m, tea.ClearScreen

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.tinput, cmd = m.tinput.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return m.header()
}

func main() {
	flag.StringVar(&icon, "icon", "", "prompt icon")
	flag.StringVar(&title, "title", "Mina", "prompt title")
	flag.Parse()

	lipgloss.SetColorProfile(termenv.TrueColor)
	p := tea.NewProgram(newMina(), tea.WithAltScreen(), tea.WithOutput(os.Stderr))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
