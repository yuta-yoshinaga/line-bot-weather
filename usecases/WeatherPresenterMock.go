package usecases

import (
	"github.com/stretchr/testify/mock"
)

// WeatherPresenterMock WeatherPresenterクラスのモック
type WeatherPresenterMock struct {
	mock.Mock
}

// NewWeatherPresenterMock コンストラクタ
func NewWeatherPresenterMock() *WeatherPresenterMock {
	weatherPresenterMock := new(WeatherPresenterMock)
	return weatherPresenterMock
}

// GetWeatherText 天気予報テキスト取得
func (mock *WeatherPresenterMock) GetWeatherText(output *WeatherOutput) string {
	return mock.Called(output).Get(0).(string)
}
