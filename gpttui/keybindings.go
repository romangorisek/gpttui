package gpttui

import "github.com/awesome-gocui/gocui"

type Keybinding struct {
	ID       string
	ViewName string
	Key      interface{}
	Modifier gocui.Modifier
	Handler  func(*gocui.Gui, *gocui.View) error
}

type KeybindingOverride struct {
	ID       string
	Key      interface{}
	Modifier gocui.Modifier
}

func (gpttui *Gpttui) SetKeybindings() error {
	keybindings := []Keybinding{
		{"global_quit", "", gocui.KeyCtrlC, gocui.ModNone, quit},
		{"global_nextView", "", gocui.KeyTab, gocui.ModNone, nextView},
	}

	// if len(gpttui.userConfig.KeybindingOverrides) > 0 {
	// 	for i := range keybindings {
	// 		for _, override := range gpttui.userConfig.KeybindingOverrides {
	// 			if keybindings[i].ID == override.ID {
	// 				keybindings[i].Key = override.Key
	// 				keybindings[i].Modifier = override.Modifier
	// 			}
	// 		}
	// 	}
	// }

	for _, kb := range keybindings {
		if err := gpttui.gui.SetKeybinding(kb.ViewName, kb.Key, kb.Modifier, kb.Handler); err != nil {
			return err
		}
	}
	return nil
}
