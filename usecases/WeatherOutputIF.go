package usecases

// WeatherOutputIF
type WeatherOutputIF interface {
	GetWeatherText(output *WeatherOutput) string
}
