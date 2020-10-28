package usecases

import "line-bot-weather/entities"

// WeatherOutput 天気情報取得アウトプットデータ
type WeatherOutput struct {
	Weather *entities.CurrentWeather
}
