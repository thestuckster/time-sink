package main

import (
	"embed"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/go-co-op/gocron"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"strconv"
	"strings"
	"time"
	"time-sink/internal"
	"time-sink/internal/bindings"
	"time-sink/internal/services"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	startup()
	timeSinkConfig := services.LoadConfiguration()

	//start scheduler
	toWatch := buildTestToWatchSet(timeSinkConfig)
	scheduler := gocron.NewScheduler(time.UTC)
	err := buildSchedule(timeSinkConfig, toWatch, scheduler)
	if err != nil {
		panic(err)
	}
	scheduler.StartAsync()

	//Create Wails app. comment all the wails lines if you just want to run the scheduler for development with no UI
	app := NewApp()

	//Create go --> js bindings and register them in the Bind array
	configBinding := bindings.NewTimeSinkConfigBinding()
	usageBinding := bindings.NewUsageBinding()
	processBinding := bindings.NewProcessBinding()

	//TODO: this is deprecated and needs deleted but doing that will break the UI and I don't feel like fixing it right this second
	// sorry future me.
	dataBinding := bindings.NewDataBinding()

	//Create application with options
	err = wails.Run(&options.App{
		Title:  "time-sink",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
			dataBinding,
			configBinding,
			processBinding,
			usageBinding,
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
	services.CreateTableIfNotExists()
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
		_, err := scheduler.Every(duration).Hours().Do(services.RecordProcesses, toWatch)
		return err
	case "m":
		_, err := scheduler.Every(duration).Minutes().Do(services.RecordProcesses, toWatch)
		return err
	case "s":
		_, err := scheduler.Every(duration).Seconds().Do(services.RecordProcesses, toWatch)
		return err
	default:
		panic("Invalid check_interval string")
	}
}
