package bindings

import (
	"context"
	"fmt"
	"time-sink/internal"
)

type ConfigDto struct {
	Applications []string `json:"applications"`

	// "1 h", "1 m", "30 s", etc. space is important!
	CheckInterval string `json:"check_interval"`
}

type TimeSinkConfigBinding struct {
	ctx context.Context
}

func NewTimeSinkConfigBinding() *TimeSinkConfigBinding { return &TimeSinkConfigBinding{} }

func (cfg *TimeSinkConfigBinding) GetConfig() ConfigDto {
	fmt.Println("&&&&& get config call")
	config := internal.LoadConfiguration()
	return ConfigDto{Applications: config.Applications, CheckInterval: config.CheckInterval}
}

// TODO: right now this requires a restart for changes to take effect. fix so it doesn't.
func (cfg *TimeSinkConfigBinding) SaveConfig(dto ConfigDto) {
	timeSinkConfig := internal.TimeSinkConfig{Applications: dto.Applications, CheckInterval: dto.CheckInterval}
	internal.SaveConfiguration(timeSinkConfig)
}
