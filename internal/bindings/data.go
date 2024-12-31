package bindings

import (
	"context"
	"time-sink/internal"
)

type ProcessUsageInfoDto struct {
	Name     string `json:"name"`
	Seen     string `json:"seen"`
	Duration string `json:"duration"`
}

type DataBinding struct {
	ctx context.Context
}

func NewDataBinding() *DataBinding { return &DataBinding{} }

// GetDailyProcesses date parameter needs format of "yyyy-MM-dd"
func (d *DataBinding) GetDailyProcesses(date string) []ProcessUsageInfoDto {
	return mapProcessesDtos(internal.GetDailyRecords(date))
}

func mapProcessesDtos(data []internal.ProcessDto) []ProcessUsageInfoDto {
	usageInfo := make([]ProcessUsageInfoDto, 0)
	for _, d := range data {
		usageInfo = append(usageInfo, ProcessUsageInfoDto{
			Name:     d.Name,
			Seen:     d.Seen,
			Duration: d.Duration,
		})
	}

	return usageInfo
}
