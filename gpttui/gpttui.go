package gpttui

import (
	"github.com/awesome-gocui/gocui"
	"github.com/romangorisek/gpttui/config"
	"github.com/romangorisek/gpttui/dataservice"
)

var views = []string{
	"question",
	"conversation",
	"sessions",
}
var activeView = 0

type Gpttui struct {
	dataService *dataservice.DataService
	gui         *gocui.Gui
	userConfig  *config.UserConfig
}

func Init(dataService *dataservice.DataService, gui *gocui.Gui, userConfig *config.UserConfig) *Gpttui {
	return &Gpttui{dataService: dataService, gui: gui, userConfig: userConfig}
}

func (gpttui *Gpttui) GocuiLayout(g *gocui.Gui) error {
	maxX, maxY := gpttui.gui.Size()
	questionHeight := 3
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen
	g.SelFrameColor = gocui.ColorGreen
	if activeView == 0 {
		g.Cursor = true
	} else {
		g.Cursor = false
	}

	if v, err := gpttui.gui.SetView(views[0], maxX/4+1, maxY-questionHeight-2, maxX-1, maxY-2, 0); err != nil {
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
	if v, err := gpttui.gui.SetView(views[1], maxX/4+1, 0, maxX-1, maxY-questionHeight-3, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Conversation"
	}
	if v, err := gpttui.gui.SetView(views[2], 0, 0, maxX/4, maxY-2, 0); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Sessions"
	}
	// TODO: add help row
	return nil
}
