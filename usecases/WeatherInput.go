package usecases

// WeatherInput 天気情報取得インプットデータ
type WeatherInput struct {
	Appid string
	City  string
	Lat   string
	Lon   string
}
