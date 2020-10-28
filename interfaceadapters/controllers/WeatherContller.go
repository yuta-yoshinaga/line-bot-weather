package controllers

import (
	"fmt"
	"line-bot-weather/usecases"
	"net/http"
)

// WeatherContller 天気情報取得コントローラ
type WeatherContller struct {
	usecase usecases.WeatherInputIF
	Appid   string
}

// NewWeatherContller コンストラクタ
func NewWeatherContller(usecase usecases.WeatherInputIF, Appid string) *WeatherContller {
	weatherContller := new(WeatherContller)
	weatherContller.usecase = usecase
	weatherContller.Appid = Appid
	return weatherContller
}

// GetWeather 天気情報取得メソッド
func (weatherContller WeatherContller) GetWeather(w http.ResponseWriter, r *http.Request) {
	input := new(usecases.WeatherInput)
	input.Appid = weatherContller.Appid
	fmt.Fprintf(w, weatherContller.usecase.GetCurrentWeather(input))
}
