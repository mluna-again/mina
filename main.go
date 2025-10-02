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
const CONFIRM_MODE = "confirm"
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
	t.Width = 80
	t.TextStyle = theme.prompt
	t.PromptStyle = theme.prompt
	t.PlaceholderStyle = theme.placeholder
	// This disables the blinking that i *didnt* enable, btw
	t.Cursor.Blur()
	t.Cursor.Style = theme.promptCursor
	t.Cursor.TextStyle = theme.promptCursor
	if mode != CONFIRM_MODE {
		t.Prompt = fmt.Sprintf("%s ", icon)
	} else {
		t.Prompt = " "
		t.Cursor.Style = theme.noCursor
		t.Cursor.TextStyle = theme.noCursor
		t.CharLimit = 3
		t.Width = 3
	}

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
	switch m.mode {
	case PROMPT_MODE:
		return m.updatePrompt(msg)
	case FZF_MODE:
		return m.updateFzf(msg)
	case CONFIRM_MODE:
		return m.updateConfirm(msg)
	default:
		panic("unknown mode")
	}
}

func (m model) View() string {
	switch m.mode {
	case PROMPT_MODE:
		return m.headerView()
	case CONFIRM_MODE:
		return m.confirmView()
	default:
		return m.headerView() + "\n" + m.listView()
	}
}

func main() {
	flag.StringVar(&icon, "icon", "", "prompt icon")
	flag.StringVar(&title, "title", "Mina", "prompt title")
	flag.StringVar(&mode, "mode", "fzf", "modes available: [prompt, fzf, confirm]")
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
