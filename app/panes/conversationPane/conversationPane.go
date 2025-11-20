package conversationPane

import (
	"log"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/romangorisek/gpttui/app/panes"
	"github.com/romangorisek/gpttui/constants"
	"github.com/romangorisek/gpttui/keys"
)

type SetInsertModeMsg struct{}

type ConversationMsgType int

const (
	Question ConversationMsgType = iota
	Answer
)

type ConversationMsg struct {
	Text     string
	ItemType ConversationMsgType
}

type Model struct {
	panes.Pane
	Messages     []ConversationMsg
	windowHeight int
	inputHeight  int
}

func New() *Model {
	return &Model{
		Pane:        panes.Pane{Title: "Conversation"},
		inputHeight: 1,
	}
}

func (m Model) View() string {
	conversationContent := m.renderMessages()
	return m.GetStyle(panes.StyleProps{}).Render(conversationContent)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width - 2
		m.windowHeight = msg.Height
		m.Height = m.windowHeight - 2 - (m.inputHeight + 2) - 2
		log.Printf("conversation height %d, input %d, term %d", m.Height, m.inputHeight, msg.Height)
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

func (m *Model) AppendQuestion(text string) {
	m.Messages = append(m.Messages, ConversationMsg{Text: text, ItemType: Question})
}

func (m *Model) SetInputHeight(height int) {
	m.inputHeight = height
	m.Height = m.windowHeight - 2 - (m.inputHeight + 2) - 2
	log.Printf("conversation height %d, input %d, term %d", m.Height, m.inputHeight, m.windowHeight)
	m.SetSize(m.Width, m.Height)
}

func (m *Model) renderMessages() string {
	bubbleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		MaxWidth(m.Width - 6).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(constants.Colors.Grey))

	rightAlign := lipgloss.NewStyle().
		Align(lipgloss.Right).
		Width(m.Width - 2)

	content := ""

	for _, message := range m.Messages {
		if message.ItemType == Question {
			bubble := bubbleStyle.Render(message.Text)
			content += rightAlign.Render(bubble) + "\n"
		}
	}

	return content
}
