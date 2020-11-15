package controllers

import (
	"io/ioutil"
	"line-bot-weather/usecases"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeatherContller_GetWeather(t *testing.T) {
	appid := "appid"
	city := "tokyo,JP"
	input := new(usecases.WeatherInput)
	input.Appid = appid
	input.City = city

	interactorMock := NewWeatherInteractorMock()
	interactorMock.On("GetCurrentWeather", input).Return("東京都の本日の天気は晴れ。現在の気温は15°、最高気温は20°、最低気温は10°です。")

	c := NewWeatherContller(interactorMock, appid)
	testserver := httptest.NewServer(http.HandlerFunc(c.GetWeather))
	defer testserver.Close()
	t.Run("success", func(t *testing.T) {
		res, err := http.Get(testserver.URL)
		assert.Equal(t, nil, err)
		resStr, err := ioutil.ReadAll(res.Body)
		defer res.Body.Close()
		assert.Equal(t, nil, err)
		assert.Equal(t, http.StatusOK, res.StatusCode)
		assert.Equal(t, "東京都の本日の天気は晴れ。現在の気温は15°、最高気温は20°、最低気温は10°です。", string(resStr))
		interactorMock.AssertCalled(t, "GetCurrentWeather", input)
	})
}
