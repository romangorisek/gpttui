package panes

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/romangorisek/gpttui/constants"
)

type paneType int

const (
	answers paneType = iota
	questionInput
)

type Pane struct {
	Title    string
	Focused  bool
	PaneType paneType
	Height   int
	Width    int
}

func (p *Pane) Focus() {
	p.Focused = true
}

func (p *Pane) Blur() {
	p.Focused = false
}

func (p *Pane) SetSize(width, height int) {
	p.Width = width
	p.Height = height
}

func (p Pane) Init() tea.Cmd {
	return nil
}

type StyleProps struct {
	NoBorder  bool
	NoPadding bool
}

func (p *Pane) GetStyle(props StyleProps) lipgloss.Style {
	style := lipgloss.NewStyle().
		Height(p.Height).
		Width(p.Width)

	if !props.NoPadding {
		style = style.Padding(1)
	}

	if !props.NoBorder {
		style = style.Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(constants.Colors.Grey))

		if p.Focused {
			style = style.
				BorderForeground(lipgloss.Color(constants.Colors.Green))
		}
	}

	return style
}
