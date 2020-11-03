package usecases

import (
	"line-bot-weather/entities"

	"github.com/stretchr/testify/mock"
)

// OpenWeatherMapRepositoryMock OpenWeatherMapRepositoryクラスのモック
type OpenWeatherMapRepositoryMock struct {
	mock.Mock
}

// NewOpenWeatherMapRepositoryMock コンストラクタ
func NewOpenWeatherMapRepositoryMock() *OpenWeatherMapRepositoryMock {
	openWeatherMapRepositoryMock := new(OpenWeatherMapRepositoryMock)
	return openWeatherMapRepositoryMock
}

// GetCurrentWeather 天気情報取得メソッド
func (mock *OpenWeatherMapRepositoryMock) GetCurrentWeather(query string) *entities.CurrentWeather {
	return mock.Called(query).Get(0).(*entities.CurrentWeather)
}
