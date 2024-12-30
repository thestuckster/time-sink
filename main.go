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

	// Keep program alive until interrupt signal is sent. this means you have to manually ctrl + c
	//sigChan := make(chan os.Signal, 1)
	//signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	//_ = <-sigChan

	// Create an instance of the app structure
	app := NewApp()

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
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

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
