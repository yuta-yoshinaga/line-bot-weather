package usecases

// WeatherOutputIF 天気情報取得アウトプットインタフェース
type WeatherOutputIF interface {
	GetWeatherText(output *WeatherOutput) string
}
