package bindings

import (
	"context"
	"time-sink/internal"
)

type ProcessUsageInfo struct {
	Name     string `json:"name"`
	Seen     string `json:"seen"`
	Duration string `json:"duration"`
}

type Data struct {
	ctx context.Context
}

func NewData() *Data { return &Data{} }

// GetDailyProcesses date parameter needs format of "yyyy-MM-dd"
func (d *Data) GetDailyProcesses(date string) []ProcessUsageInfo {
	return mapProcessesDtos(internal.GetDailyRecords(date))
}

func mapProcessesDtos(data []internal.ProcessDto) []ProcessUsageInfo {
	usageInfo := make([]ProcessUsageInfo, 0)
	for _, d := range data {
		usageInfo = append(usageInfo, ProcessUsageInfo{
			Name:     d.Name,
			Seen:     d.Seen,
			Duration: d.Duration,
		})
	}

	return usageInfo
}
