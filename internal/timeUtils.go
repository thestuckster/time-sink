package internal

import "time"

func ToRealDate(t time.Time) float64 {
	return float64(t.UnixNano() / 1e9)
}

func FromRealDate(unixEpch float64) time.Time {
	return time.Unix(0, int64(unixEpch*1e9))
}

func ToStandardDateFormat(date time.Time) string {
	return date.Format("2006-01-02")
}
