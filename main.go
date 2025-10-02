package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

const PROMPT_MODE = "prompt"
const CONFIRM_MODE = "confirm"
const FZF_MODE = "fzf"
const MENU_MODE = "menu"

var response string

// FLAGS
var icon string
var title string
var mode string
var separator string
var displayColumns string
var height int
var width int

type model struct {
	theme        theme
	mode         string
	width        int
	height       int
	ignoreHeight bool
	ignoreWidth  bool
	tinput       textinput.Model
	list         list.Model
	title        string
	content      []string
}

func newMina() model {
	switch mode {
	case FZF_MODE:
		return newFzfModel()
	case PROMPT_MODE:
		return newPromptModel()
	case CONFIRM_MODE:
		return newConfirmModel()
	case MENU_MODE:
		return newMenuModel()
	default:
		panic("unknown mode")
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
	case MENU_MODE:
		return m.updateMenu(msg)
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
	case MENU_MODE:
		return m.menuView()
	default:
		return m.headerView() + "\n" + m.listView()
	}
}

func main() {
	flag.StringVar(&icon, "icon", "", "prompt icon")
	flag.StringVar(&title, "title", "Mina", "prompt title")
	flag.StringVar(&mode, "mode", "fzf", "modes available: [prompt, fzf, confirm, menu]")
	flag.StringVar(&separator, "sep", " ", "separator used with -nth")
	flag.StringVar(&displayColumns, "nth", "", "display specific columns. eg: -nth 1 displays only the second column, -nth 0,3 displays 1st, 2nd and 3rd column.")
	flag.IntVar(&height, "height", 0, "height, if 0 or empty it takes the full screen")
	flag.IntVar(&width, "width", 0, "width, if 0 or empty it takes the full screen")
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
