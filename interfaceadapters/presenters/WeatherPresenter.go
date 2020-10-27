package presenters

import (
	"line-bot-weather/usecases"
	"strconv"
)

// WeatherPresenter
type WeatherPresenter struct {
}

// NewWeatherPresenter コンストラクタ
func NewWeatherPresenter() *WeatherPresenter {
	weatherPresenter := new(WeatherPresenter)
	return weatherPresenter
}

// GetWeatherText 天気予報テキスト取得
func (weatherPresenter WeatherPresenter) GetWeatherText(output *usecases.WeatherOutput) string {
	return output.Weather.Name + "の本日の天気は" + output.Weather.Weather[0].Description + "。最高気温は" + strconv.FormatFloat(output.Weather.Main.TempMax, 'f', -1, 64) + "°、最低気温は" + strconv.FormatFloat(output.Weather.Main.TempMin, 'f', -1, 64) + "°です。"
}
