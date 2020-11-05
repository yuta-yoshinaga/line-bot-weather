package presenters

import (
	"line-bot-weather/entities"
	"line-bot-weather/usecases"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetCurrentWeather1 GetCurrentWeatherのテスト
func TestWeatherPresenter_GetWeatherText(t *testing.T) {
	t.Run("success GetCurrentWeather()", func(t *testing.T) {
		weatherPresenter := NewWeatherPresenter()
		currentWeather := &usecases.WeatherOutput{
			Weather: &entities.CurrentWeather{
				Name: "東京都",
				Weather: []entities.Weather{
					{Description: "晴れ"},
				},
				Main: entities.Main{
					TempMax: 22.2,
					TempMin: 11.1,
				},
			},
		}
		res := weatherPresenter.GetWeatherText(currentWeather)
		assert.Equal(t, "東京都の本日の天気は晴れ。最高気温は22.2°、最低気温は11.1°です。", res)
	})
}
