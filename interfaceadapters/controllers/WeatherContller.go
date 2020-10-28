package controllers

import (
	"fmt"
	"line-bot-weather/usecases"
	"net/http"
	"os"
)

// WeatherContller 天気情報取得コントローラ
type WeatherContller struct {
	usecase usecases.WeatherInputIF
}

// NewWeatherContller コンストラクタ
func NewWeatherContller(usecase usecases.WeatherInputIF) *WeatherContller {
	weatherContller := new(WeatherContller)
	weatherContller.usecase = usecase
	return weatherContller
}

// GetWeather 天気情報取得メソッド
func (weatherContller WeatherContller) GetWeather(w http.ResponseWriter, r *http.Request) {
	input := new(usecases.WeatherInput)
	input.Appid = os.Getenv("APPID")
	fmt.Fprintf(w, weatherContller.usecase.GetCurrentWeather(input))
}
