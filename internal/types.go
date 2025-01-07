package internal

import (
	"time"
)

type Process struct {
	Id       int
	Name     string
	Seen     string // mm/DD/yyyy
	Duration string //HH:MM:SS
	Time     time.Time
}

type ProcessUsageDbDto struct {
	Id       int
	Name     string
	Seen     string // mm/DD/yyyy
	Duration string //HH:MM:SS
}

type TimeSinkConfig struct {
	Applications []string `json:"applications"`

	// "1 h", "1 m", "30 s", etc. space is important!
	CheckInterval string `json:"check_interval"`
}
