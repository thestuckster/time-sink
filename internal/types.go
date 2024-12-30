package internal

import "time"

type Process struct {
	Id       int
	Name     string
	Seen     string // mm/DD/yyyy
	Duration string //HH:MM:SS
	Time     time.Time
}

type ProcessDto struct {
	Id       int
	Name     string
	Seen     string // mm/DD/yyyy
	Duration string //HH:MM:SS
}
