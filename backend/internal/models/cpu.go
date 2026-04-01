type CPUMetrics struct {
    // From cpu.Percent()
    // calls cpu.Percent(0, false)
    OverallPercent  float64   `json:"overall_percent"`
    PerCorePercent  []float64 `json:"per_core_percent"`

    // From type TimeStats struct
    // call cpu.TimesStat and returns String()
    User    float64 `json:"user"`
    System  float64 `json:"system"`
    Idle    float64 `json:"idle"`
    Iowait  float64 `json:"iowait"`
    Steal   float64 `json:"steal"`

    // From type InfoState struct
    // call cpu.InfoStat which returns a string
    ModelName string  `json:"model_name"`
    Cores     int32   `json:"cores"`
    Mhz       float64 `json:"mhz"`
    CacheSize int32 `json"CacheSize"`

    // Metadata
    Timestamp int64 `json:"timestamp"` // unix ms, for charting x-axis
}