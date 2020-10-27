package entities

// Wind Wind
type Wind struct {
	Speed     float64 `json:"speed"`      // Wind speed. Unit Default: meter/sec, Metric: meter/sec, Imperial: miles/hour.
	Deg       int     `json:"deg"`        // Wind direction, degrees (meteorological)
	GrndLevel int     `json:"grnd_level"` // Atmospheric pressure on the ground level, hPa
}
