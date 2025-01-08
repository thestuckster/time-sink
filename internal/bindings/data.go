package bindings

import (
	"context"
	"time-sink/internal"
	"time-sink/internal/repository"
)

type ProcessUsageData struct {
	Name     string `json:"name"`
	Seen     int64  `json:"seen"`
	Duration int64  `json:"duration"`
}

type DataBinding struct {
	ctx context.Context
}

func NewDataBinding() *DataBinding { return &DataBinding{} }

// GetDailyProcesses date parameter needs format of "yyyy-MM-dd"
func (d *DataBinding) GetDailyProcesses() []ProcessUsageData {
	return mapProcessesDtos(internal.GetDailyRecords())
}

func mapProcessesDtos(data []repository.ApplicationDto) []ProcessUsageData {
	usageInfo := make([]ProcessUsageData, 0)
	for _, d := range data {

		usageInfo = append(usageInfo, ProcessUsageData{
			Name:     d.Name,
			Seen:     d.Seen,
			Duration: d.Duration,
		})
	}

	return usageInfo
}
