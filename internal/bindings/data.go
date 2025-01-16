package bindings

import (
	"context"
	"time"
	"time-sink/internal/repository"
	"time-sink/internal/services"
)

type ProcessUsageData struct {
	Name     string `json:"name"`
	Seen     int64  `json:"seen"`
	Duration int64  `json:"duration"`
}

type TotalDuration struct {
	TotalDuration int64 `json:"total_duration"`
}

type DataBinding struct {
	ctx context.Context
}

func NewDataBinding() *DataBinding { return &DataBinding{} }

func (d *DataBinding) GetDailyProcesses() []ProcessUsageData {
	return mapAllToDtos(services.GetDailyRecords())
}

func (d *DataBinding) GetLast30Days() map[string]TotalDuration {
	today := time.Now()
	past := today.AddDate(0, 0, -30)

	applicationRecords := services.GetRecordsInRange(past, today)

	response := make(map[string]TotalDuration)
	for _, record := range applicationRecords {
		if val, ok := response[record.Name]; ok {
			val.TotalDuration += record.Duration
		} else {
			response[record.Name] = TotalDuration{
				TotalDuration: record.Duration,
			}
		}
	}

	return response
}

func mapAllToDtos(applicationRecords []repository.ApplicationRecordEntity) []ProcessUsageData {
	usageInfo := make([]ProcessUsageData, 0)
	for _, d := range applicationRecords {

		usageInfo = append(usageInfo, ProcessUsageData{
			Name:     d.Name,
			Seen:     d.Seen,
			Duration: d.Duration,
		})
	}

	return usageInfo
}
