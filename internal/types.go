package internal

import (
	"time"
)

type Process struct {
	Id       int
	Name     string
	Seen     time.Time
	Duration int64     // delta of Seen and now, in minutes
	Time     time.Time //TODO: remove if not needed
}

type ProcessUsageDbDto struct {
	Id       int
	Name     string
	Seen     float64 //nanoseconds of unix epoch
	Duration int64   // delta of Seen and now, in minutes
}

type TimeSinkConfig struct {
	Applications []string `json:"applications"`

	// "1 h", "1 m", "30 s", etc. space is important!
	CheckInterval string `json:"check_interval"`
}
