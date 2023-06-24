package presenters

import (
	"testing"

	"github.com/yuta-yoshinaga/line-bot-weather/entities"
	"github.com/yuta-yoshinaga/line-bot-weather/usecases"

	"github.com/stretchr/testify/assert"
)

// TestWeatherPresenter_GetWeatherText GetWeatherTextのテスト
func TestWeatherPresenter_GetWeatherText(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		weatherPresenter := NewWeatherPresenter()
		currentWeather := &usecases.WeatherOutput{
			Weather: &entities.CurrentWeather{
				Name: "東京都",
				Weather: []entities.Weather{
					{Description: "晴れ"},
				},
				Main: entities.Main{
					Temp:    12.3,
					TempMax: 22.2,
					TempMin: 11.1,
				},
			},
		}
		res := weatherPresenter.GetWeatherText(currentWeather)
		assert.Equal(t, "東京都の本日の天気は晴れ。現在の気温は12.3°、最高気温は22.2°、最低気温は11.1°です。", res)
	})
	t.Run("failed", func(t *testing.T) {
		weatherPresenter := NewWeatherPresenter()
		res := weatherPresenter.GetWeatherText(nil)
		assert.Equal(t, "", res)
	})
	t.Run("failed", func(t *testing.T) {
		weatherPresenter := NewWeatherPresenter()
		currentWeather := &usecases.WeatherOutput{
			Weather: nil,
		}
		res := weatherPresenter.GetWeatherText(currentWeather)
		assert.Equal(t, "", res)
	})
	t.Run("failed", func(t *testing.T) {
		weatherPresenter := NewWeatherPresenter()
		currentWeather := &usecases.WeatherOutput{
			Weather: &entities.CurrentWeather{
				Weather: []entities.Weather{},
			},
		}
		res := weatherPresenter.GetWeatherText(currentWeather)
		assert.Equal(t, "", res)
	})
}
