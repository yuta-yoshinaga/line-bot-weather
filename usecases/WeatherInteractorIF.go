package usecases

// WeatherInteractorIF 天気情報取得インプットインタフェース
type WeatherInteractorIF interface {
	GetCurrentWeather(input *WeatherInput) string
}
