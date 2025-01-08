package internal

import (
	"time"
)

type Process struct {
	Id       int
	Name     string
	Seen     time.Time
	Duration int64 // delta of Seen and now
}

type TimeSinkConfig struct {
	Applications []string `json:"applications"`

	// "1 h", "1 m", "30 s", etc. space is important!
	CheckInterval string `json:"check_interval"`
}
