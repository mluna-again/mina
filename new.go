package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

func newFzfModel() model {
	theme := getTheme()
	t := newInput(theme)

	content := []string{}
	var err error
	content, err = loadInput()
	if err != nil {
		panic(err)
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
	l.SetHeight(height)
	l.SetWidth(width - theme.listItem.GetPaddingLeft() - theme.listItem.GetPaddingRight())

	ignoreH := height != 0
	ignoreW := width != 0

	return model{
		tinput:       t,
		title:        title,
		theme:        theme,
		mode:         mode,
		content:      content,
		list:         l,
		height:       height,
		width:        width,
		ignoreHeight: ignoreH,
		ignoreWidth:  ignoreW,
	}
}

func newPromptModel() model {
	theme := getTheme()
	t := newInput(theme)
	ignoreH := height != 0
	ignoreW := width != 0

	return model{
		tinput:       t,
		title:        title,
		theme:        theme,
		mode:         mode,
		height:       height,
		width:        width,
		ignoreHeight: ignoreH,
		ignoreWidth:  ignoreW,
	}
}

func newConfirmModel() model {
	theme := getTheme()
	t := newInput(theme)
	ignoreH := height != 0
	ignoreW := width != 0

	return model{
		tinput:       t,
		title:        title,
		theme:        theme,
		mode:         mode,
		height:       height,
		width:        width,
		ignoreHeight: ignoreH,
		ignoreWidth:  ignoreW,
	}
}

func newMenuModel() model {
	theme := getTheme()
	ignoreH := height != 0
	ignoreW := width != 0

	return model{
		title:        title,
		theme:        theme,
		mode:         mode,
		height:       height,
		width:        width,
		ignoreHeight: ignoreH,
		ignoreWidth:  ignoreW,
	}
}

func newInput(theme theme) textinput.Model {
	t := textinput.New()
	t.Focus()
	t.TextStyle = theme.prompt
	t.PromptStyle = theme.prompt
	t.PlaceholderStyle = theme.placeholder
	// This disables the blinking that i *didnt* enable, btw
	t.Cursor.Blur()
	t.Cursor.Style = theme.promptCursor
	t.Cursor.TextStyle = theme.promptCursor
	t.Width = width - lipgloss.Width(t.Prompt) - 1
	if mode != CONFIRM_MODE {
		t.Prompt = fmt.Sprintf("%s ", icon)
	} else {
		t.Prompt = " "
		t.Cursor.Style = theme.noCursor
		t.Cursor.TextStyle = theme.noCursor
		t.CharLimit = 3
		t.Width = 3
	}

	return t
}
