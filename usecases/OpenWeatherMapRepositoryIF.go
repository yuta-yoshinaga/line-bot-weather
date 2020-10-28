package usecases

import "line-bot-weather/entities"

// OpenWeatherMapRepositoryIF 天気情報取得リポジトリインタフェース
type OpenWeatherMapRepositoryIF interface {
	GetCurrentWeather(query string) *entities.CurrentWeather
}
