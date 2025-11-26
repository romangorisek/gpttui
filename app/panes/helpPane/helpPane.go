package helpPane

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/romangorisek/gpttui/app/panes"
	"github.com/romangorisek/gpttui/keys"
)

type Model struct {
	panes.Pane
	help help.Model
}

func New() *Model {
	helpModel := help.New()
	helpModel.ShowAll = false

	return &Model{
		Pane: panes.Pane{Title: "Help", Height: 2},
		help: helpModel,
	}
}

func (m Model) View() string {
	helpView := m.help.View(keys.Keys)
	return m.GetStyle(panes.StyleProps{NoBorder: true}).Padding(0, 1).Render(helpView)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	return m, cmd
}
