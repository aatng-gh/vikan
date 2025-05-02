package main

import "github.com/charmbracelet/lipgloss"

var (
	DocStyle     = lipgloss.NewStyle().Margin(1, 2).Border(lipgloss.NormalBorder())
	FocusedStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
	NormalStyle  = lipgloss.NewStyle().Border(lipgloss.HiddenBorder())
)
