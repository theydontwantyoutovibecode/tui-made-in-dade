package main

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.help.SetWidth(msg.Width)
		m.list.SetSize(msg.Width-6, msg.Height-14)
		m.viewport.SetWidth(msg.Width - 8)
		m.viewport.SetHeight(msg.Height - 8)
		return m, nil

	case tea.KeyMsg:
		if m.state == viewTextInput {
			return m.updateTextInput(msg)
		}

		switch {
		case key.Matches(msg, m.keys.Quit):
			if m.state == viewWelcome {
				m.quitting = true
				return m, tea.Quit
			}
			m.state = viewWelcome
			return m, nil
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, nil
		case key.Matches(msg, m.keys.Back):
			if m.state != viewWelcome {
				m.state = viewWelcome
				m.input.Reset()
			}
			return m, nil
		}
	}

	switch m.state {
	case viewWelcome:
		return m.updateWelcome(msg)
	case viewViewport:
		return m.updateViewport(msg)
	default:
		return m, nil
	}
}

func (m model) updateWelcome(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	if keyMsg, ok := msg.(tea.KeyMsg); ok && key.Matches(keyMsg, m.keys.Enter) {
		if item, ok := m.list.SelectedItem().(menuItem); ok {
			m.state = item.target
			if item.target == viewTextInput {
				cmd := m.input.Focus()
				return m, cmd
			}
		}
	}

	return m, cmd
}

func (m model) updateTextInput(msg tea.Msg) (tea.Model, tea.Cmd) {
	if keyMsg, ok := msg.(tea.KeyMsg); ok {
		if key.Matches(keyMsg, m.keys.Back) {
			m.state = viewWelcome
			m.input.Reset()
			return m, nil
		}
		k := keyMsg.Key()
		if k.Code == 'c' && k.Mod.Contains(tea.ModCtrl) {
			m.state = viewWelcome
			m.input.Reset()
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m model) updateViewport(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}
