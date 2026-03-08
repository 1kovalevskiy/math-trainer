package ui

import "github.com/charmbracelet/lipgloss"

var (
	Panel = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("69")).
		Background(lipgloss.Color("237")).
		Padding(1, 2)

	Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("230")).
		Background(lipgloss.Color("62")).
		Padding(0, 1)

	Subtitle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("117"))

	Label = lipgloss.NewStyle().
		Foreground(lipgloss.Color("159"))

	Value = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("230"))

	Hint = lipgloss.NewStyle().
		Foreground(lipgloss.Color("246"))

	Error = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("203"))

	Accent = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("221"))

	MenuActive = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("230")).
			Background(lipgloss.Color("62")).
			Padding(0, 1)

	MenuInactive = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252")).
			Padding(0, 1)
)

func MenuItem(active bool, text string) string {
	if active {
		return MenuActive.Render("▸ " + text)
	}

	return MenuInactive.Render("  " + text)
}
