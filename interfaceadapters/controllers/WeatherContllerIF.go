package controllers

import "net/http"

// WeatherContllerIF 天気情報取得コントローラインタフェース
type WeatherContllerIF interface {
	GetWeather(w http.ResponseWriter, r *http.Request)
}
