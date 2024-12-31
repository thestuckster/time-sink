package main

import (
	"embed"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/go-co-op/gocron"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"time"
	"time-sink/internal"
	"time-sink/internal/bindings"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	startup()

	// TODO: this is just looking for firefox and discord. need a GUI to really make this useful
	toWatch := buildTestToWatchSet()
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Every(1).Minutes().Do(internal.RecordProcesses, toWatch)
	if err != nil {
		panic(err)
	}
	scheduler.StartAsync()

	// Create Wails app. comment if you just want to run the scheduler for development with no UI
	app := NewApp()

	//Create go --> js bindings and register them in the Bind array
	data := bindings.NewData()

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "time-sink",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			data,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

	// wait for manual ctrl+c interrupt. uncomment if you don't want to start all of wails
	//sigChan := make(chan os.Signal, 1)
	//signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	//_ = <-sigChan

}

func startup() {
	internal.CreateDbFile()
	internal.CreateTableIfNotExists()
}

// TODO: delete after ensuring things work
func buildTestToWatchSet() *hashset.Set {
	set := hashset.New()
	set.Add("firefox.exe")
	set.Add("Discord.exe")
	set.Add("Marvel-Win64-Shipping.exe")

	return set
}
