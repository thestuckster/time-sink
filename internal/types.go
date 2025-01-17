package internal

type TimeSinkConfig struct {
	Applications []string `json:"applications"`

	// "1 h", "1 m", "30 s", etc. space is important!
	CheckInterval string `json:"check_interval"`
}
