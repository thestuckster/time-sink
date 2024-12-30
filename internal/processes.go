package internal

import (
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/mitchellh/go-ps"
	"time"
)

func RecordProcesses(toWatch *hashset.Set) {
	fmt.Println("Recording Processes")
	procs, err := ps.Processes()
	if err != nil {
		panic(err)
	}

	// because you can have multiple processes for the same applications, Chrome / Firefox are good examples.
	// this ensures we only track them once.
	alreadyRecorded := hashset.New()
	for _, proc := range procs {
		name := proc.Executable()
		if toWatch.Contains(name) && !alreadyRecorded.Contains(name) {
			SaveSeenProcess(Process{
				Name: name,
				Seen: time.Now(),
			})
			alreadyRecorded.Add(name)
		}
	}
}
