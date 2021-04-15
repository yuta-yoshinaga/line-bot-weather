package controllers

import (
	"line-bot-weather/usecases"

	"github.com/stretchr/testify/mock"
)

// WeatherInteractorMock WeatherInteractorクラスのモック
type WeatherInteractorMock struct {
	mock.Mock
}

// NewWeatherInteractorMock コンストラクタ
func NewWeatherInteractorMock() *WeatherInteractorMock {
	weatherInteractorMock := new(WeatherInteractorMock)
	return weatherInteractorMock
}

// GetCurrentWeather 天気情報取得メソッド
func (mock *WeatherInteractorMock) GetCurrentWeather(input *usecases.WeatherInput) string {
	return mock.Called(input).Get(0).(string)
}
