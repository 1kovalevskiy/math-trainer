package result

import (
	"github.com/1kovalevskiy/math-trainer/internal/app/tui/ui"
	mathmodels "github.com/1kovalevskiy/math-trainer/internal/models/math"
	"github.com/charmbracelet/lipgloss"
)

type Option int

const (
	OptionRetry Option = iota
	OptionSettings
	OptionHome
)

type Model struct {
	summary mathmodels.TrainingSummary
	cursor  int
	options []string

	scrollOffset       int
	lastViewportHeight int
	lastContentRows    int
	viewportWidth      int
	viewportHeight     int
}

func NewModel() Model {
	return Model{
		cursor: 0,
		options: []string{
			"Решить еще один пример",
			"Настройки сложности",
			"В главное меню",
		},
	}
}

func (m Model) WithSummary(summary *mathmodels.TrainingSummary) Model {
	if summary != nil {
		m.summary = *summary
	}
	m.cursor = 0
	m.scrollOffset = 0
	m.lastViewportHeight = 0
	m.lastContentRows = 0
	m.viewportWidth = 0
	m.viewportHeight = 0
	return m
}

func (m Model) selectedOption() Option {
	if m.cursor < 0 || m.cursor >= len(m.options) {
		return OptionRetry
	}

	return Option(m.cursor)
}

func (m *Model) clampScroll() {
	m.scrollOffset = m.scrollState().ClampOffset(m.scrollOffset)
}

func (m Model) maxScrollOffset() int {
	return m.scrollState().MaxOffset()
}

func (m Model) WithViewport(width int, height int) Model {
	if width < 1 {
		width = 1
	}
	if height < 1 {
		height = 1
	}
	m.viewportWidth = width
	m.viewportHeight = height
	return m
}

func (m *Model) refreshScrollBounds() {
	viewportHeight := m.viewportHeight
	if viewportHeight < 1 {
		viewportHeight = 1
	}
	m.lastViewportHeight = viewportHeight

	if len(m.summary.Results) == 0 {
		m.lastContentRows = 1
		m.clampScroll()
		return
	}

	entryWidth := 1
	for _, entry := range m.summary.Results {
		entryWidth = max(entryWidth, lipgloss.Width(renderEntry(entry)))
	}

	width := m.viewportWidth
	if width < 1 {
		width = 1
	}
	layout := m.gridLayout(width, viewportHeight, len(m.summary.Results), entryWidth)
	m.lastContentRows = layout.ContentRows
	m.clampScroll()
}

func (m Model) scrollState() ui.ScrollState {
	return ui.ScrollState{
		Offset:       m.scrollOffset,
		ViewportRows: m.lastViewportHeight,
		ContentRows:  m.lastContentRows,
	}
}
