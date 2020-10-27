package usecases

import "line-bot-weather/entities"

// OpenWeatherMapRepositoryIF
type OpenWeatherMapRepositoryIF interface {
	GetCurrentWeather(query string) *entities.CurrentWeather
}
