package bindings

import (
	"log"
	"maps"
	"slices"
	"sort"
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

func (usb *UsageBinding) GetAllTimeUsage() []UsageInfo {
	allApps := services.GetAll()

	totals := make(map[string]UsageInfo)
	for _, app := range allApps {
		if value, ok := totals[app.Name]; ok {
			temp := value
			temp.Duration += app.Duration
			totals[app.Name] = temp
		} else {
			totals[app.Name] = UsageInfo{app.Name, app.Duration}
		}
	}

	summedApps := slices.Collect(maps.Values(totals))
	sort.Slice(summedApps, func(i, j int) bool {
		return summedApps[i].Duration > summedApps[j].Duration
	})
	return summedApps
}
