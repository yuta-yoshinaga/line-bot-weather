package entities

// CurrentWeather 天気予報構造体　https://openweathermap.org/current
type CurrentWeather struct {
	Coord      Coord     `json:"coord"`      // City geo location
	Weather    []Weather `json:"weather"`    // more info Weather condition codes)
	Base       string    `json:"base"`       // Internal parameter
	Main       Main      `json:"main"`       // Temperature
	Visibility int       `json:"visibility"` // Visibility, meter
	Wind       Wind      `json:"wind"`       // Wind
	Clouds     Clouds    `json:"clouds"`     // Clouds
	Rain       Rain      `json:"rain"`       // Rain
	Snow       Rain      `json:"snow"`       // Snow
	Dt         int       `json:"dt"`         // Time of data calculation, unix, UTC
	Sys        Sys       `json:"sys"`        // System
	Timezone   int       `json:"timezone"`   // Shift in seconds from UTC
	ID         int       `json:"id"`         // City ID
	Name       string    `json:"name"`       // City name
	Cod        int       `json:"cod"`        // Internal parameter
}
