package usecases

// WeatherInputIF 天気情報取得インプットインタフェース
type WeatherInputIF interface {
	GetCurrentWeather(input *WeatherInput) string
}
