package entities

// Rain Rain
type Rain struct {
	M1h float64 `json:"1h"` // Rain volume for the last 1 hour, mm
	M3h float64 `json:"3h"` // Rain volume for the last 3 hours, mm
}
