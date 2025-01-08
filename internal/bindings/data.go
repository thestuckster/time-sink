package bindings

import (
	"context"
	"time-sink/internal"
)

type ProcessUsageData struct {
	Name     string `json:"name"`
	Seen     string `json:"seen"`
	Duration int64  `json:"duration"`
}

type DataBinding struct {
	ctx context.Context
}

func NewDataBinding() *DataBinding { return &DataBinding{} }

// GetDailyProcesses date parameter needs format of "yyyy-MM-dd"
func (d *DataBinding) GetDailyProcesses(date string) []ProcessUsageData {
	return mapProcessesDtos(internal.GetDailyRecords(date))
}

func mapProcessesDtos(data []internal.ApplicationDto) []ProcessUsageData {
	usageInfo := make([]ProcessUsageData, 0)
	for _, d := range data {

		internal.ToStandardDateFormat(internal.FromRealDate(d.Seen))
		usageInfo = append(usageInfo, ProcessUsageData{
			Name:     d.Name,
			Seen:     internal.ToStandardDateFormat(internal.FromRealDate(d.Seen)),
			Duration: d.Duration,
		})
	}

	return usageInfo
}
