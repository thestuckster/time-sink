package main

import (
	"embed"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/go-co-op/gocron"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
	"time-sink/internal"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	startup()
	timeSinkConfig := internal.LoadConfiguration()

	//start scheduler
	toWatch := buildTestToWatchSet(timeSinkConfig)
	scheduler := gocron.NewScheduler(time.UTC)
	err := buildSchedule(timeSinkConfig, toWatch, scheduler)
	if err != nil {
		panic(err)
	}
	scheduler.StartAsync()

	//Create Wails app. comment if you just want to run the scheduler for development with no UI
	//app := NewApp()
	//
	////Create go --> js bindings and register them in the Bind array
	//dataBinding := bindings.NewDataBinding()
	//configBinding := bindings.NewTimeSinkConfigBinding()
	//processBinding := bindings.NewProcessBinding()
	//
	////Create application with options
	//err = wails.Run(&options.App{
	//	Title:  "time-sink",
	//	Width:  1024,
	//	Height: 768,
	//	AssetServer: &assetserver.Options{
	//		Assets: assets,
	//	},
	//	BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
	//	OnStartup:        app.startup,
	//	Bind: []interface{}{
	//		app,
	//		dataBinding,
	//		configBinding,
	//		processBinding,
	//	},
	//})
	//
	//if err != nil {
	//	println("Error:", err.Error())
	//}

	// wait for manual ctrl+c interrupt. uncomment if you don't want to start all of wails
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	_ = <-sigChan

}

func startup() {
	internal.CreateDbFile()
	internal.CreateTableIfNotExists()
}

func buildTestToWatchSet(config internal.TimeSinkConfig) *hashset.Set {
	set := hashset.New()
	for _, app := range config.Applications {
		set.Add(app)
	}

	return set
}

func buildSchedule(config internal.TimeSinkConfig, toWatch *hashset.Set, scheduler *gocron.Scheduler) error {
	configuredInterval := config.CheckInterval
	parts := strings.Split(configuredInterval, " ")

	if len(parts) != 2 {
		panic("Invalid check_interval string")
	}

	duration, err := strconv.Atoi(parts[0])
	if err != nil {
		panic("invalid check_interval duration")
	}

	switch parts[1] {
	case "h":
		_, err := scheduler.Every(duration).Hours().Do(internal.RecordProcesses, toWatch)
		return err
	case "m":
		_, err := scheduler.Every(duration).Minutes().Do(internal.RecordProcesses, toWatch)
		return err
	case "s":
		_, err := scheduler.Every(duration).Seconds().Do(internal.RecordProcesses, toWatch)
		return err
	default:
		panic("Invalid check_interval string")
	}
}
