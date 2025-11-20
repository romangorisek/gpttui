package inputPane

import (
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/romangorisek/gpttui/app/panes"
	"github.com/romangorisek/gpttui/constants"
	"github.com/romangorisek/gpttui/keys"
)

type Mode int

const maxHeight = 4
const (
	Normal Mode = iota
	Insert
	Visual
)

type Model struct {
	panes.Pane
	Mode          Mode
	InputText     string
	Textarea      textarea.Model
	prevLineCount int
}

type InputSubmittedMsg struct {
	Text string
}

type PaneResizeMsg struct {
	Height int
}

func New() *Model {
	textarea := textarea.New()
	textarea.ShowLineNumbers = false
	textarea.SetHeight(1)
	textarea.Cursor.SetMode(cursor.CursorBlink)
	textarea.FocusedStyle.CursorLine = lipgloss.NewStyle()
	textarea.Prompt = ""
	textarea.KeyMap.InsertNewline.SetKeys("ctrl+n")

	return &Model{
		Pane:     panes.Pane{Title: "Question input"},
		Textarea: textarea,
	}
}

func (m *Model) SetMode(mode Mode) {
	m.Mode = mode

	if mode == Insert {
		m.Textarea.Focus()
	} else if m.Textarea.Focused() {
		m.Textarea.Blur()
	}
}

func (m *Model) SetSize(width, height int) {
	m.Height = height
	m.Width = width
	m.Textarea.SetWidth(width - 2)
}

func (m Model) View() string {
	style := m.GetStyle(panes.StyleProps{})
	if m.Mode == Insert {
		style = style.
			BorderForeground(lipgloss.Color(constants.Colors.Yellow))

	}
	style = style.Padding(0, 1)
	return style.Render(m.Textarea.View())
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

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
	cmds = append(cmds, cmd)

	return m, cmd
}

func (m *Model) handleNormalModeKeyMsgs(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, keys.Keys.InsertMode):
		m.SetMode(Insert)
	}
	return m, nil
}

func (m *Model) handleInsertModeKeyMsgs(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch {
	case msg.Type == tea.KeyEsc:
		m.SetMode(Normal)
	case msg.Type == tea.KeyEnter:
		text := m.Textarea.Value()
		m.Textarea.SetValue("")
		resizeCmd := m.resize(1)
		m.SetMode(Normal)

		submittedCmd := func() tea.Msg {
			return InputSubmittedMsg{Text: text}
		}

		cmds = append(cmds, resizeCmd, submittedCmd)
	default:
		m.Textarea, cmd = m.Textarea.Update(msg)
		cmds = append(cmds, cmd)

		newHeight := calcHeight(m.Textarea.Value(), m.Textarea.Width())
		if newHeight != m.Textarea.Height() {
			resizeCmd := m.resize(newHeight)
			cmds = append(cmds, resizeCmd)
		}
	}

	return m, tea.Batch(cmds...)

}

func (m *Model) resize(newHeight int) tea.Cmd {
	m.Textarea.SetHeight(newHeight)
	// TODO: when we resize up we have one line invisible and one empty, need to scroll
	// TODO: paste from clipboard is broken
	return func() tea.Msg {
		return PaneResizeMsg{Height: newHeight}
	}
}

func (m *Model) handleVisualModeKeyMsgs(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	return m, nil
}

func calcHeight(text string, width int) int {
	newHeight := 0
	for _, line := range strings.Split(text, "\n") {
		// we add the soft breaks to the line count
		newHeight += 1 + (len(line) / width)
	}
	if newHeight < 1 {
		newHeight = 1
	}
	if newHeight > maxHeight {
		newHeight = maxHeight
	}

	return newHeight
}
