package helpPane

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/romangorisek/gpttui/app/panes"
)

type Model struct {
	panes.Pane
}

func New() *Model {
	// help := help.New()
	// help.ShowAll = true

	return &Model{
		Pane: panes.Pane{Title: "Help", Height: 2},
	}
}

func (m Model) View() string {
	return m.GetStyle(panes.StyleProps{NoBorder: true}).Padding(0, 1).Render("some help line stuff")
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	return m, cmd
}
