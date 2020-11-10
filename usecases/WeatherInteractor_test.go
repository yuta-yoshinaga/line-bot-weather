package usecases

import (
	"line-bot-weather/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestWeatherInteractor_GetCurrentWeather GetCurrentWeatherのテスト
func TestWeatherInteractor_GetCurrentWeather(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo := NewOpenWeatherMapRepositoryMock()
		mockRepo.On("GetCurrentWeather", "&appid=&lang=ja&units=metric").Return(new(entities.CurrentWeather)).Once()

		testWeatherOut := new(WeatherOutput)
		testWeatherOut.Weather = new(entities.CurrentWeather)
		mockPre := NewWeatherPresenterMock()
		mockPre.On("GetWeatherText", testWeatherOut).Return("東京都の本日の天気は晴れ。現在の気温は15°、最高気温は20°、最低気温は10°です。").Once()

		weatherInteractor := NewWeatherInteractor(mockRepo, mockPre)

		input := new(WeatherInput)
		currentWeather := weatherInteractor.GetCurrentWeather(input)
		assert.Equal(t, "東京都の本日の天気は晴れ。現在の気温は15°、最高気温は20°、最低気温は10°です。", currentWeather)
		mockRepo.AssertCalled(t, "GetCurrentWeather", "&appid=&lang=ja&units=metric")
		mockRepo.AssertNumberOfCalls(t, "GetCurrentWeather", 1)
		mockPre.AssertCalled(t, "GetWeatherText", testWeatherOut)
		mockPre.AssertNumberOfCalls(t, "GetWeatherText", 1)
	})
	t.Run("success", func(t *testing.T) {
		mockRepo := NewOpenWeatherMapRepositoryMock()
		mockRepo.On("GetCurrentWeather", "q=City&lat=Lat&lon=Lon&appid=Appid&lang=ja&units=metric").Return(new(entities.CurrentWeather)).Once()

		testWeatherOut := new(WeatherOutput)
		testWeatherOut.Weather = new(entities.CurrentWeather)
		mockPre := NewWeatherPresenterMock()
		mockPre.On("GetWeatherText", testWeatherOut).Return("東京都の本日の天気は晴れ。現在の気温は15°、最高気温は20°、最低気温は10°です。").Once()

		weatherInteractor := NewWeatherInteractor(mockRepo, mockPre)

		input := new(WeatherInput)
		input.Appid = "Appid"
		input.City = "City"
		input.Lat = "Lat"
		input.Lon = "Lon"
		currentWeather := weatherInteractor.GetCurrentWeather(input)
		assert.Equal(t, "東京都の本日の天気は晴れ。現在の気温は15°、最高気温は20°、最低気温は10°です。", currentWeather)
		mockRepo.AssertCalled(t, "GetCurrentWeather", "q=City&lat=Lat&lon=Lon&appid=Appid&lang=ja&units=metric")
		mockRepo.AssertNumberOfCalls(t, "GetCurrentWeather", 1)
		mockPre.AssertCalled(t, "GetWeatherText", testWeatherOut)
		mockPre.AssertNumberOfCalls(t, "GetWeatherText", 1)
	})
}
