package main

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m model) View() tea.View {
	if m.quitting {
		return tea.NewView("")
	}

	var s string
	switch m.state {
	case viewTextInput:
		s = m.viewTextInput()
	case viewViewport:
		s = m.viewViewport()
	case viewDevTips:
		s = m.viewDevTips()
	default:
		s = m.viewWelcome()
	}

	v := tea.NewView(s)
	v.AltScreen = true
	return v
}

func (m model) viewWelcome() string {
	var b strings.Builder

	header := headerBoxStyle.Width(min(m.width-4, 60)).Render(
		titleStyle.Render("myapp v0.0.1") + "\n" +
			subtitleStyle.Render("A TUI application built with Bubbletea"),
	)
	b.WriteString(header)
	b.WriteString("\n\n")
	b.WriteString("  Welcome! This is your new TUI application.\n\n")

	b.WriteString(m.list.View())
	b.WriteString("\n\n")

	helpView := m.help.View(m.keys)
	b.WriteString(helpStyle.Width(m.width).Render(helpView))

	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, b.String())
}

func (m model) viewTextInput() string {
	var b strings.Builder

	header := inputLabelStyle.Render("Text Input Demo")
	b.WriteString(header)
	b.WriteString("\n\n")
	b.WriteString("Type something and see it echoed below.\n\n")
	b.WriteString(m.input.View())
	b.WriteString("\n\n")

	if m.input.Value() != "" {
		b.WriteString(echoStyle.Render(fmt.Sprintf("You typed: %s", m.input.Value())))
	} else {
		b.WriteString(subtitleStyle.Render("Start typing..."))
	}

	b.WriteString("\n\n")
	b.WriteString(subtitleStyle.Render("Press esc to go back"))

	box := contentBoxStyle.Width(min(m.width-4, 60)).Render(b.String())
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)
}

func (m model) viewViewport() string {
	header := inputLabelStyle.Render("Viewport Demo") + "  " +
		subtitleStyle.Render("(scroll with arrow keys, PgUp/PgDn)")

	footer := subtitleStyle.Render("Press esc to go back")

	content := header + "\n\n" + m.viewport.View() + "\n\n" + footer
	box := contentBoxStyle.Width(min(m.width-4, 70)).Render(content)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)
}

func (m model) viewDevTips() string {
	tips := fmt.Sprintf(`%s

%s
  model.go   — Application state (what data you track)
  update.go  — Handle messages (user input, events)
  view.go    — Render the UI (returns a string)
  keys.go    — Keybinding definitions
  styles.go  — Lipgloss style definitions

%s
  Messages are immutable events (key press, timer, etc.)
  Update returns a new model + optional commands
  Commands produce side effects (I/O, timers, HTTP)
  View is a pure function: model → string

%s
  Use tea.LogToFile() for debug logging
  Run with DEBUG=1 to enable file logging
  Components manage their own state (embed in model)
  AltScreen mode is set via View.AltScreen = true

%s`,
		tipHeaderStyle.Render("TUI Development Tips"),
		tipHeaderStyle.Render("File Structure:"),
		tipHeaderStyle.Render("Elm Architecture:"),
		tipHeaderStyle.Render("Debugging:"),
		subtitleStyle.Render("Press esc to go back"),
	)

	box := contentBoxStyle.Width(min(m.width-4, 65)).Render(tips)
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)
}
