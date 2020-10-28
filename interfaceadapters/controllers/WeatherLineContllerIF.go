package controllers

import "net/http"

// WeatherLineContllerIF 天気情報取得コントローラインタフェース
type WeatherLineContllerIF interface {
	GetWeatherLine(w http.ResponseWriter, r *http.Request)
}
