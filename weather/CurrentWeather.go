package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"line-bot-weather/entities"
	"net/http"
	"strconv"
)

// CurrentWeather 天気予報構造体　https://openweathermap.org/current
type CurrentWeather struct {
	Coord      entities.Coord     `json:"coord"`      // City geo location
	Weather    []entities.Weather `json:"weather"`    // more info Weather condition codes)
	Base       string             `json:"base"`       // Internal parameter
	Main       entities.Main      `json:"main"`       // Temperature
	Visibility int                `json:"visibility"` // Visibility, meter
	Wind       entities.Wind      `json:"wind"`       // Wind
	Clouds     entities.Clouds    `json:"clouds"`     // Clouds
	Rain       entities.Rain      `json:"rain"`       // Rain
	Snow       entities.Rain      `json:"snow"`       // Snow
	Dt         int                `json:"dt"`         // Time of data calculation, unix, UTC
	Sys        entities.Sys       `json:"sys"`        // System
	Timezone   int                `json:"timezone"`   // Shift in seconds from UTC
	ID         int                `json:"id"`         // City ID
	Name       string             `json:"name"`       // City name
	Cod        int                `json:"cod"`        // Internal parameter
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
