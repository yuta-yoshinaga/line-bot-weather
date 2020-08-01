package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Coord City geo location
type Coord struct {
	Lon float64 `json:"lon"` // longitude
	Lat float64 `json:"lat"` // latitude
}

// Weather more info Weather condition codes)
type Weather struct {
	ID          int    `json:"id"`          // Weather condition id
	Main        string `json:"main"`        // Group of weather parameters (Rain, Snow, Extreme etc.)
	Description string `json:"description"` // Weather condition within the group. You can get the output in your language. Learn more
	Icon        string `json:"icon"`        // Weather icon id
}

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

// Wind Wind
type Wind struct {
	Speed     float64 `json:"speed"`      // Wind speed. Unit Default: meter/sec, Metric: meter/sec, Imperial: miles/hour.
	Deg       int     `json:"deg"`        // Wind direction, degrees (meteorological)
	GrndLevel int     `json:"grnd_level"` // Atmospheric pressure on the ground level, hPa
}

// Clouds Clouds
type Clouds struct {
	All int `json:"all"` // Cloudiness, %
}

// Rain Rain
type Rain struct {
	M1h float64 `json:"1h"` // Rain volume for the last 1 hour, mm
	M3h float64 `json:"3h"` // Rain volume for the last 3 hours, mm
}

// Sys System
type Sys struct {
	Type    int     `json:"type"`    // Internal parameter
	ID      int     `json:"id"`      // Internal parameter
	Message float64 `json:"message"` // Internal parameter
	Country string  `json:"country"` // Country code (GB, JP etc.)
	Sunrise int     `json:"sunrise"` // Sunrise time, unix, UTC
	Sunset  int     `json:"sunset"`  // Sunset time, unix, UTC
}

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

// NewCurrentWeather コンストラクタ
func NewCurrentWeather(query string) *CurrentWeather {
	weather := new(CurrentWeather)
	url := "https://api.openweathermap.org/data/2.5/weather"
	url += "?" + query
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
	jsonBytes := ([]byte)(byteArray)
	if err := json.Unmarshal(jsonBytes, weather); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return nil
	}
	return weather
}

// GetWeatherText 天気予報テキスト取得
func (currentWeather CurrentWeather) GetWeatherText() string {
	return currentWeather.Name + "の本日の天気は" + currentWeather.Weather[0].Description + "。最高気温は" + strconv.FormatFloat(currentWeather.Main.TempMax, 'f', -1, 64) + "°、最低気温は" + strconv.FormatFloat(currentWeather.Main.TempMin, 'f', -1, 64) + "°です。"
}
