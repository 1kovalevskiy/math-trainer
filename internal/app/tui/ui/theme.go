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

	buttonActive = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("230")).
			Background(lipgloss.Color("62")).
			Padding(0, 2)

	buttonInactive = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252")).
			Background(lipgloss.Color("238")).
			Padding(0, 2)

	smallButtonActive = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("230")).
				Background(lipgloss.Color("99")).
				Padding(0, 1)

	smallButtonInactive = lipgloss.NewStyle().
				Foreground(lipgloss.Color("252")).
				Background(lipgloss.Color("239")).
				Padding(0, 1)
)

func MenuItem(active bool, text string) string {
	return Button(text, active)
}

func Button(text string, active bool) string {
	if active {
		return buttonActive.Render(text)
	}

	return buttonInactive.Render(text)
}

func SmallButton(text string, active bool) string {
	if active {
		return smallButtonActive.Render(text)
	}

	return smallButtonInactive.Render(text)
}
