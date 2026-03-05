package main

import "charm.land/lipgloss/v2"

var (
	primaryColor   = lipgloss.Color("#FF79C6")
	secondaryColor = lipgloss.Color("#50FA7B")
	subtleColor    = lipgloss.Color("#6272A4")
	textColor      = lipgloss.Color("#F8F8F2")

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(primaryColor).
			Align(lipgloss.Center)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(subtleColor).
			Align(lipgloss.Center)

	headerBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(primaryColor).
			Padding(0, 2).
			Align(lipgloss.Center)

	contentBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(secondaryColor).
			Padding(1, 2)

	helpStyle = lipgloss.NewStyle().
			Foreground(subtleColor).
			Align(lipgloss.Center)

	tipHeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#F1FA8C"))

	tipTextStyle = lipgloss.NewStyle().
			Foreground(textColor)

	inputLabelStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(secondaryColor)

	echoStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Italic(true)
)
