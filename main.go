package main

import (
	"github.com/awesome-gocui/gocui"
	"github.com/romangorisek/gpttui/config"
	"github.com/romangorisek/gpttui/gpttui"
	"github.com/romangorisek/gpttui/llmservice"
	"github.com/romangorisek/gpttui/logger"
)

func main() {
	log := logger.New()

	gui, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panic(err)
	}
	defer gui.Close()

	config := config.LoadUserConfig()

	API_KEY := "jakjsdflkajsdfkjasdflkj" // TODO: actually read this form the env - think of a sensible default to read (SOMETHING_API_KEY) or it sould be customisable from config

	llmService, err := llmservice.New(log, "GPT_4", API_KEY)
	if err != nil {
		log.Panicln(err)
	}

	llmService.Test()

	app := gpttui.Init(llmService, gui, config)
	gui.SetManagerFunc(app.GocuiLayout)

	if err := app.SetKeybindings(); err != nil {
		log.Panicln(err)
	}

	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
