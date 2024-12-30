package main

import (
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/go-co-op/gocron"
	"os"
	"os/signal"
	"syscall"
	"time"
	"time-sink/internal"
)

func main() {
	startup()

	// TODO: this is just looking for firefox and discord. need a GUI to really make this useful
	toWatch := buildTestToWatchSet()
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Every(1).Minutes().Do(internal.RecordProcesses, toWatch)
	if err != nil {
		return
	}
	scheduler.StartAsync()

	// Keep program alive until interrupt signal is sent. this means you have to manually ctrl + c
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	_ = <-sigChan
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

	return set
}
