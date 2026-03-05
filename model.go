package main

import (
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/textinput"
	"charm.land/bubbles/v2/viewport"
	tea "charm.land/bubbletea/v2"
)

type viewState int

const (
	viewWelcome viewState = iota
	viewTextInput
	viewViewport
	viewDevTips
)

type model struct {
	state    viewState
	list     list.Model
	help     help.Model
	keys     keyMap
	input    textinput.Model
	viewport viewport.Model
	width    int
	height   int
	quitting bool
}

type menuItem struct {
	title       string
	description string
	target      viewState
}

func (i menuItem) Title() string       { return i.title }
func (i menuItem) Description() string { return i.description }
func (i menuItem) FilterValue() string { return i.title }

func initialModel() model {
	items := []list.Item{
		menuItem{"Text Input Demo", "Try the textinput component", viewTextInput},
		menuItem{"Viewport Demo", "Scrollable content area", viewViewport},
		menuItem{"View Dev Tips", "TUI development guidance", viewDevTips},
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "What would you like to explore?"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.SetShowHelp(false)

	ti := textinput.New()
	ti.Placeholder = "Type something..."
	ti.CharLimit = 256

	vp := viewport.New()
	vp.SetContent(viewportContent())

	return model{
		state:    viewWelcome,
		list:     l,
		help:     help.New(),
		keys:     defaultKeyMap(),
		input:    ti,
		viewport: vp,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func viewportContent() string {
	return `This is a viewport demo. It displays scrollable content.

The viewport component from Bubbles handles:
  - Scrolling with arrow keys, Page Up/Down, Home/End
  - Mouse wheel scrolling
  - Content that exceeds the visible area

You can use viewports for:
  - Log viewers
  - Document readers
  - Help text panels
  - Any content that needs scrolling

Lorem ipsum dolor sit amet, consectetur adipiscing elit.
Sed do eiusmod tempor incididunt ut labore et dolore magna
aliqua. Ut enim ad minim veniam, quis nostrud exercitation
ullamco laboris nisi ut aliquip ex ea commodo consequat.

Duis aute irure dolor in reprehenderit in voluptate velit
esse cillum dolore eu fugiat nulla pariatur. Excepteur sint
occaecat cupidatat non proident, sunt in culpa qui officia
deserunt mollit anim id est laborum.

Scroll down to see more content...

This demonstrates how the viewport handles overflow content.
The scroll position indicator appears on the right side.

You can also set the viewport content dynamically by calling
vp.SetContent() with new text at any time.

Press 'esc' to go back to the menu.`
}
