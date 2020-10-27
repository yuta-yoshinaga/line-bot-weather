package entities

// Main Temperature
type Main struct {
	Temp      float64 `json:"temp"`       // Temperature. Unit Default: Kelvin, Metric: Celsius, Imperial: Fahrenheit.
	FeelsLike float64 `json:"feels_like"` // Temperature. This temperature parameter accounts for the human perception of weather. Unit Default: Kelvin, Metric: Celsius, Imperial: Fahrenheit.
	Pressure  int     `json:"pressure"`   // Atmospheric pressure (on the sea level, if there is no sea_level or grnd_level data), hPa
	Humidity  int     `json:"humidity"`   // Humidity, %
	TempMin   float64 `json:"temp_min"`   // Minimum temperature at the moment. This is minimal currently observed temperature (within large megalopolises and urban areas). Unit Default: Kelvin, Metric: Celsius, Imperial: Fahrenheit.
	TempMax   float64 `json:"temp_max"`   // Maximum temperature at the moment. This is maximal currently observed temperature (within large megalopolises and urban areas). Unit Default: Kelvin, Metric: Celsius, Imperial: Fahrenheit.
	SeaLevel  int     `json:"sea_level"`  // Atmospheric pressure on the sea level, hPa
	GrndLevel int     `json:"grnd_level"` // Atmospheric pressure on the ground level, hPa
}
