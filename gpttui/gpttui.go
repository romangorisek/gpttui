package gpttui

import (
	"fmt"

	"github.com/awesome-gocui/gocui"
	"github.com/romangorisek/gpttui/config"
	"github.com/romangorisek/gpttui/llmservice"
)

var views = []string{
	"question",
	"conversation",
	"help",
}
var activeView = 0

type Gpttui struct {
	llmService *llmservice.LlmService
	gui        *gocui.Gui
	userConfig *config.UserConfig
}

func Init(llmService *llmservice.LlmService, gui *gocui.Gui, userConfig *config.UserConfig) *Gpttui {
	return &Gpttui{llmService: llmService, gui: gui, userConfig: userConfig}
}

func (gpttui *Gpttui) GocuiLayout(g *gocui.Gui) error {
	maxX, maxY := gpttui.gui.Size()
	questionHeight := 3
	helpHeight := 2
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen
	g.SelFrameColor = gocui.ColorGreen
	if activeView == 0 {
		g.Cursor = true
	} else {
		g.Cursor = false
	}

	if v, err := gpttui.gui.SetView(views[0], 0, maxY-questionHeight-helpHeight-1, maxX-1, maxY-helpHeight-1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Question"
		v.Editable = true
		v.Wrap = true

		if _, err = setCurrentViewOnTop(g, views[0]); err != nil {
			return err
		}
	}
	if v, err := gpttui.gui.SetView(views[1], 0, 0, maxX-1, maxY-questionHeight-helpHeight-2, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Conversation"
	}
	if v, err := gpttui.gui.SetView(views[2], 0, maxY-helpHeight-1, maxX-1, maxY-1, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Frame = false
		fmt.Fprint(v, "some help page")
	}
	return nil
}
