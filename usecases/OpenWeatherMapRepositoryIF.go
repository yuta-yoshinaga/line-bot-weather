package usecases

import "github.com/yuta-yoshinaga/line-bot-weather/entities"

// OpenWeatherMapRepositoryIF 天気情報取得リポジトリインタフェース
type OpenWeatherMapRepositoryIF interface {
	GetCurrentWeather(query string) *entities.CurrentWeather
}
