package conversationPane

import (
	"log"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/romangorisek/gpttui/app/panes"
	"github.com/romangorisek/gpttui/keys"
)

type SetInsertModeMsg struct{}

type Model struct {
	panes.Pane
}

func New() *Model {
	return &Model{
		Pane: panes.Pane{Title: "Conversation"},
	}
}

func (m Model) View() string {
	return m.GetStyle(panes.StyleProps{}).Render()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width - 2
		m.Height = msg.Height - 2 - 3 - 2 // TODO: needs to be reactive
		log.Printf("conversation height %d", m.Height)
		m.SetSize(m.Width, m.Height)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.InsertMode):
			return m, func() tea.Msg {
				return SetInsertModeMsg{}
			}
		}
	}
	return m, cmd
}
