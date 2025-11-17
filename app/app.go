package app

import (
	"log"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/romangorisek/gpttui/app/panes/conversationPane"
	"github.com/romangorisek/gpttui/app/panes/helpPane"
	"github.com/romangorisek/gpttui/app/panes/inputPane"
	"github.com/romangorisek/gpttui/keys"
)

/*
 NOTE:
 - should I rename InputPane struct to Model (and same in other panes)?
 - maybe move RemoveLastWord function from InputPane to some utils strings.go or something?

 TODO:
 - add text area component to the inputPane
 - add appendQuestion function in conversationPane
   - should display a nice bubble in with the text
 - chat gpt service to call api (copy from the old project)
 - display actual help with with help bubbletea component
 - display api response (if not error) in the conversation
 - handle error - show a pop-up with the message
 - make the height of the conversationPane be reactive to changes in height of the inputPane


*/

type App struct {
	focused          int
	loaded           bool
	quitting         bool
	helpPane         *helpPane.Model
	conversationPane *conversationPane.Model
	inputPane        *inputPane.Model
}

func New() error {
	helpPane := helpPane.New()
	conversationPane := conversationPane.New()
	inputPane := inputPane.New()
	inputPane.Focus()
	app := &App{helpPane: helpPane, focused: 0, conversationPane: conversationPane, inputPane: inputPane}

	p := tea.NewProgram(app, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}

func (app *App) Init() tea.Cmd {
	return nil
}

func (app *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		app.helpPane.Update(msg)
		app.conversationPane.Update(msg)
		app.inputPane.Update(msg)
		app.loaded = true
		return app, nil

	case inputPane.InputSubmittedMsg:
		// app.conversationPane.appendQuestion(msg.Text)
		log.Printf("submit the message: %s", msg.Text)
	case conversationPane.SetInsertModeMsg:
		app.conversationPane.Blur()
		app.inputPane.Focus()
		app.inputPane.SetMode(inputPane.Insert)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Keys.GlobalQuit):
			app.quitting = true
			return app, tea.Quit
		case key.Matches(msg, keys.Keys.Quit):
			if !(app.inputPane.Focused && app.inputPane.Mode == inputPane.Insert) {
				app.quitting = true
				return app, tea.Quit
			}
		case key.Matches(msg, keys.Keys.Next):
			if app.inputPane.Focused {
				app.inputPane.Blur()
				app.conversationPane.Focus()
			} else {
				app.inputPane.Focus()
				app.conversationPane.Blur()

			}
		default:
			if app.inputPane.Focused {
				newInputPane, cmd := app.inputPane.Update(msg)
				app.inputPane = newInputPane.(*inputPane.Model)
				cmds = append(cmds, cmd)
			}
			if app.conversationPane.Focused {
				newConversationPane, cmd := app.conversationPane.Update(msg)
				app.conversationPane = newConversationPane.(*conversationPane.Model)
				cmds = append(cmds, cmd)
			}
		}
	}
	return app, tea.Batch(cmds...)
}

func (app *App) View() string {
	if app.quitting {
		return ""
	}
	if !app.loaded {
		return "loading..."
	}

	return lipgloss.JoinVertical(
		lipgloss.Top,
		app.conversationPane.View(),
		app.inputPane.View(),
		app.helpPane.View(),
	)
}
