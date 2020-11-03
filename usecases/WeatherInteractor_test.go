package usecases

import (
	"line-bot-weather/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetCurrentWeather2 GetCurrentWeatherのテスト
func TestGetCurrentWeather2(t *testing.T) {
	mockRepo := NewOpenWeatherMapRepositoryMock()
	mockRepo.On("GetCurrentWeather", "&appid=&lang=ja&units=metric").Return(new(entities.CurrentWeather)).Once()

	testWeatherOut := new(WeatherOutput)
	testWeatherOut.Weather = new(entities.CurrentWeather)
	mockPre := NewWeatherPresenterMock()
	mockPre.On("GetWeatherText", testWeatherOut).Return("東京都の本日の天気は晴れ。最高気温は20°、最低気温は10°です。").Once()

	weatherInteractor := NewWeatherInteractor(mockRepo, mockPre)
	t.Run("success NewWeatherInteractor()", func(t *testing.T) {
		if weatherInteractor == nil {
			t.Errorf("NG")
		}
	})
	t.Run("success GetCurrentWeather()", func(t *testing.T) {
		input := new(WeatherInput)
		currentWeather := weatherInteractor.GetCurrentWeather(input)
		if currentWeather == "" {
			t.Errorf("NG")
			return
		}
		assert.Equal(t, "東京都の本日の天気は晴れ。最高気温は20°、最低気温は10°です。", currentWeather)
	})
	t.Run("success GetCurrentWeather()", func(t *testing.T) {
		mockRepo := NewOpenWeatherMapRepositoryMock()
		mockRepo.On("GetCurrentWeather", "q=City&lat=Lat&lon=Lon&appid=Appid&lang=ja&units=metric").Return(new(entities.CurrentWeather)).Once()

		testWeatherOut := new(WeatherOutput)
		testWeatherOut.Weather = new(entities.CurrentWeather)
		mockPre := NewWeatherPresenterMock()
		mockPre.On("GetWeatherText", testWeatherOut).Return("東京都の本日の天気は晴れ。最高気温は20°、最低気温は10°です。").Once()

		weatherInteractor := NewWeatherInteractor(mockRepo, mockPre)

		input := new(WeatherInput)
		input.Appid = "Appid"
		input.City = "City"
		input.Lat = "Lat"
		input.Lon = "Lon"
		currentWeather := weatherInteractor.GetCurrentWeather(input)
		if currentWeather == "" {
			t.Errorf("NG")
			return
		}
		assert.Equal(t, "東京都の本日の天気は晴れ。最高気温は20°、最低気温は10°です。", currentWeather)
	})
}
