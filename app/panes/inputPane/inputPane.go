package inputPane

import (
	"log"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/romangorisek/gpttui/app/panes"
	"github.com/romangorisek/gpttui/constants"
	"github.com/romangorisek/gpttui/keys"
)

type Mode int

const (
	Normal Mode = iota
	Insert
	Visual
)

type Model struct {
	panes.Pane
	Mode      Mode
	InputText string
}

type InputSubmittedMsg struct {
	Text string
}

func New() *Model {
	return &Model{
		Pane: panes.Pane{Title: "Question input"},
	}
}

func (m *Model) SetMode(mode Mode) {
	m.Mode = mode
}

func (m Model) View() string {
	style := m.GetStyle(panes.StyleProps{})
	if m.Mode == Insert {
		style = style.
			BorderForeground(lipgloss.Color(constants.Colors.Yellow))

	}
	style = style.Padding(0, 1)
	return style.Render(m.InputText)
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width - 2
		m.Height = 1

		m.SetSize(m.Width, m.Height)
	case tea.KeyMsg:
		switch m.Mode {
		case Insert:
			return m.handleInsertModeKeyMsgs(msg)
		case Visual:
			return m.handleVisualModeKeyMsgs(msg)
		default:
			return m.handleNormalModeKeyMsgs(msg)
		}
	}
	return m, cmd
}

// TODO: roman - change here to not compare string but to read key bindings from keys.go
func (m *Model) handleNormalModeKeyMsgs(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, keys.Keys.InsertMode):
		log.Printf("set insert mode")
		m.SetMode(Insert)
	}
	return m, nil
}

func (m *Model) handleInsertModeKeyMsgs(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case msg.Type == tea.KeyEsc:
		log.Printf("set normal mode")
		m.SetMode(Normal)
	case msg.Type == tea.KeyEnter:
		log.Printf("submit the query... %s", m.InputText)
		text := m.InputText
		m.InputText = ""
		m.SetMode(Normal)
		return m, func() tea.Msg {
			return InputSubmittedMsg{Text: text}
		}
	case msg.String() == "ctrl+w":
		m.InputText = RemoveLastWord(m.InputText)
	case (len(msg.Runes) == 1 && msg.Type == tea.KeyRunes) || msg.Type == tea.KeySpace:
		// log.Printf("adding to input text: %s", msg.String())
		m.InputText = m.InputText + msg.String()
	case msg.Type == tea.KeyCtrlN:
		m.InputText = m.InputText + "\n"
	case msg.Type == tea.KeyBackspace:
		_, size := utf8.DecodeLastRuneInString(m.InputText)
		m.InputText = m.InputText[:len(m.InputText)-size]
	}

	return m, nil
}

func (m *Model) handleVisualModeKeyMsgs(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	return m, nil
}

func RemoveLastWord(s string) string {
	s = strings.TrimRight(s, " ")

	if s == "" {
		return ""
	}

	idx := strings.LastIndex(s, " ")
	if idx == -1 {
		return ""
	}

	return s[:idx+1]
}
