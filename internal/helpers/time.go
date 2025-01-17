package helpers

import "time"

func GetStartOfDayUnixEpoch(t time.Time) int64 {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
}
