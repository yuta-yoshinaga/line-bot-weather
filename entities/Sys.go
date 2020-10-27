package entities

// Sys System
type Sys struct {
	Type    int     `json:"type"`    // Internal parameter
	ID      int     `json:"id"`      // Internal parameter
	Message float64 `json:"message"` // Internal parameter
	Country string  `json:"country"` // Country code (GB, JP etc.)
	Sunrise int     `json:"sunrise"` // Sunrise time, unix, UTC
	Sunset  int     `json:"sunset"`  // Sunset time, unix, UTC
}
