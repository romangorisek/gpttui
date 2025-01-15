package main

import (
	"log"

	"github.com/awesome-gocui/gocui"
	"github.com/romangorisek/gpttui/config"
	"github.com/romangorisek/gpttui/dataservice"
	"github.com/romangorisek/gpttui/dbservice"
	"github.com/romangorisek/gpttui/gpttui"
	"github.com/romangorisek/gpttui/logger"
	"github.com/romangorisek/gpttui/networkservice"
)

func main() {
	gui, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	defer gui.Close()

	logger.InitLogger()

	sqliteService := dbservice.NewSqliteService("asdfkafd")
	chatgptService := networkservice.NewChatgptService()

	dataService := dataservice.NewDataService(sqliteService, chatgptService)

	config := config.LoadUserConfig()

	app := gpttui.Init(dataService, gui, config)
	gui.SetManagerFunc(app.GocuiLayout)

	if err := app.SetKeybindings(); err != nil {
		log.Panicln(err)
	}

	if err := gui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
