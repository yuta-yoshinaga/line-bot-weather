package usecases

// WeatherPresenterIF 天気情報取得アウトプットインタフェース
type WeatherPresenterIF interface {
	GetWeatherText(output *WeatherOutput) string
}
