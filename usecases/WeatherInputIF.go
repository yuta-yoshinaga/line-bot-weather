package usecases

// WeatherInputIF
type WeatherInputIF interface {
	GetCurrentWeather(input *WeatherInput) string
}
