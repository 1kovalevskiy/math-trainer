package ui

import "github.com/charmbracelet/lipgloss"

const MinPanelContentWidth = 28

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
		Padding(0, 2)

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
			Align(lipgloss.Center).
			Padding(0, 2)

	buttonInactive = lipgloss.NewStyle().
			Foreground(lipgloss.Color("252")).
			Background(lipgloss.Color("238")).
			Align(lipgloss.Center).
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

	settingRowActive = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("230")).
				Background(lipgloss.Color("62"))

	settingRowInactive = lipgloss.NewStyle().
				Foreground(lipgloss.Color("252")).
				Background(lipgloss.Color("236"))

	settingRowMarker = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("221")).
				Background(lipgloss.Color("62"))
)

func MenuItem(active bool, text string) string {
	return Button(text, active)
}

func MenuItemFixed(active bool, text string, width int) string {
	return ButtonFixed(text, active, width)
}

func Button(text string, active bool) string {
	if active {
		return buttonActive.Render(text)
	}

	return buttonInactive.Render(text)
}

func ButtonFixed(text string, active bool, width int) string {
	if width < 1 {
		width = 1
	}
	if active {
		return buttonActive.Width(width).Render(text)
	}

	return buttonInactive.Width(width).Render(text)
}

func SmallButton(text string, active bool) string {
	if active {
		return smallButtonActive.Render(text)
	}

	return smallButtonInactive.Render(text)
}

func SettingRowStyle(active bool) lipgloss.Style {
	if active {
		return settingRowActive
	}

	return settingRowInactive
}

func SettingRowMarkerStyle() lipgloss.Style {
	return settingRowMarker
}
