package gpttui

import (
	"github.com/awesome-gocui/gocui"
)

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (activeView + 1) % len(views)
	name := views[nextIndex]

	if _, err := setCurrentViewOnTop(g, name); err != nil {
		return err
	}

	activeView = nextIndex
	return nil
}

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}
