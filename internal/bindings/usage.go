package bindings

import (
	"log"
	"time"
	"time-sink/internal/services"
)

type UsageBinding struct {
}

type UsageInfo struct {
	Name     string
	Duration int64
}

func NewUsageBinding() *UsageBinding {
	return &UsageBinding{}
}

func (usb *UsageBinding) GetUsageBetweenDates(start, end time.Time) []UsageInfo {
	log.Printf("DEBUG: GetUsageBetweenDates start:%v end:%v", start, end)
	response := make([]UsageInfo, 0)

	dbRecords := services.GetAllApplicationsBetweenDates(start, end)
	for _, dbRecord := range dbRecords {
		response = append(response, UsageInfo{dbRecord.Name, dbRecord.Duration})
	}

	return response
}
