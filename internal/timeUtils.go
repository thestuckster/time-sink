package internal

import "time"

func ToStandardDateFormat(date time.Time) string {
	return date.Format("2006-01-02")
}

func FormatTimeToHHMMSS(t time.Time) string {
	return t.Format("15:04:05")
}

func ParseHHMMSS(durationTimeStamp string) time.Time {
	t, err := time.Parse("15:04:05", durationTimeStamp)
	if err != nil {
		panic(err)
	}

	return t
}

func ParseStandardDateFormat(date string) time.Time {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}

	return t
}

func StartingDuration() string {
	return "00:00:00"
}
