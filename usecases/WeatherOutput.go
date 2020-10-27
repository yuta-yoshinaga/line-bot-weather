package usecases

import "line-bot-weather/entities"

// WeatherOutput
type WeatherOutput struct {
	Weather *entities.CurrentWeather
}
