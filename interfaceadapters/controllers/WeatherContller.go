package controllers

import (
	"fmt"
	"net/http"

	"github.com/yuta-yoshinaga/line-bot-weather/usecases"
)

// WeatherContller 天気情報取得コントローラ
type WeatherContller struct {
	usecase usecases.WeatherInteractorIF
	Appid   string
}

// NewWeatherContller コンストラクタ
func NewWeatherContller(usecase usecases.WeatherInteractorIF, Appid string) *WeatherContller {
	weatherContller := new(WeatherContller)
	weatherContller.usecase = usecase
	weatherContller.Appid = Appid
	return weatherContller
}

// GetWeather 天気情報取得メソッド
func (weatherContller WeatherContller) GetWeather(w http.ResponseWriter, r *http.Request) {
	input := new(usecases.WeatherInput)
	input.Appid = weatherContller.Appid
	input.City = "tokyo,JP"
	fmt.Fprintf(w, "%s", weatherContller.usecase.GetCurrentWeather(input))
}
