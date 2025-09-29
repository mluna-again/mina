package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

const PROMPT_MODE = "prompt"
const FZF_MODE = "fzf"

var response string

// FLAGS
var icon string
var title string
var mode string
var separator string
var displayColumns string

type model struct {
	theme   theme
	mode    string
	width   int
	height  int
	tinput  textinput.Model
	list    list.Model
	title   string
	content []string
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

	content := []string{}
	if mode == FZF_MODE {
		var err error
		content, err = loadInput()
		if err != nil {
			panic(err)
		}
	}

	items := []list.Item{}
	for _, line := range content {
		items = append(items, item{text: line, style: theme.listItem, selectedStyle: theme.selectedListItem})
	}

	l := list.New(items, itemDelegate{}, 0, 0)
	l.SetFilteringEnabled(false)
	l.SetShowFilter(false)
	l.SetShowHelp(false)
	l.SetShowPagination(false)
	l.SetShowStatusBar(false)
	l.SetShowTitle(false)
	l.KeyMap.CursorDown = key.NewBinding(key.WithKeys("ctrl+n"))
	l.KeyMap.CursorUp = key.NewBinding(key.WithKeys("ctrl+p"))

	return model{
		tinput:  t,
		title:   title,
		theme:   theme,
		mode:    mode,
		content: content,
		list:    l,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	m.tinput, cmd = m.tinput.Update(msg)
	cmds = append(cmds, cmd)

	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		m.tinput.Width = m.calculatePromptWidth()
		m.list.SetHeight(m.height - 3) // header
		m.list.SetWidth(m.width)
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
			response = m.content[m.list.GlobalIndex()]
			return m, tea.Quit

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.mode == PROMPT_MODE {
		return m.header()
	}

	list := m.theme.list.Render(m.list.View())
	list = lipgloss.Place(m.list.Width(), m.list.Height(), lipgloss.Center,lipgloss.Center, list, lipgloss.WithWhitespaceBackground(m.theme.list.GetBackground()))
	return m.header() + list
}

func main() {
	flag.StringVar(&icon, "icon", "", "prompt icon")
	flag.StringVar(&title, "title", "Mina", "prompt title")
	flag.StringVar(&mode, "mode", "fzf", "modes available: [prompt, fzf]")
	flag.StringVar(&separator, "sep", " ", "separator used with -nth")
	flag.StringVar(&displayColumns, "nth", "", "display specific columns. eg: -nth 1 displays only the second column, -nth 0,3 displays 1st, 2nd and 3rd column.")
	flag.Parse()

	lipgloss.SetColorProfile(termenv.TrueColor)
	p := tea.NewProgram(newMina(), tea.WithAltScreen(), tea.WithOutput(os.Stderr))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	if response != "" {
		fmt.Println(response)
	}
}
