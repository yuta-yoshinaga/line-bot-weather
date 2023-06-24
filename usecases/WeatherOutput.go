package usecases

import "github.com/yuta-yoshinaga/line-bot-weather/entities"

// WeatherOutput 天気情報取得アウトプットデータ
type WeatherOutput struct {
	Weather *entities.CurrentWeather
}
