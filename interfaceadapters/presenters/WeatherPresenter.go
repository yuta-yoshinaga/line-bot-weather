package presenters

import (
	"fmt"
	"strconv"

	"github.com/yuta-yoshinaga/line-bot-weather/usecases"
)

// WeatherPresenter 天気情報取得プレゼンター
type WeatherPresenter struct {
}

// NewWeatherPresenter コンストラクタ
func NewWeatherPresenter() *WeatherPresenter {
	weatherPresenter := new(WeatherPresenter)
	return weatherPresenter
}

// GetWeatherText 天気予報テキスト取得
func (weatherPresenter WeatherPresenter) GetWeatherText(output *usecases.WeatherOutput) string {
	if err := weatherPresenter.validate(output); err != nil {
		return ""
	}
	return output.Weather.Name + "の本日の天気は" + output.Weather.Weather[0].Description + "。現在の気温は" + strconv.FormatFloat(output.Weather.Main.Temp, 'f', -1, 64) + "°、最高気温は" + strconv.FormatFloat(output.Weather.Main.TempMax, 'f', -1, 64) + "°、最低気温は" + strconv.FormatFloat(output.Weather.Main.TempMin, 'f', -1, 64) + "°です。"
}

// validate 入力オブジェクト検証
func (weatherPresenter WeatherPresenter) validate(output *usecases.WeatherOutput) error {
	if output == nil {
		return fmt.Errorf("output nil")
	}
	if output.Weather == nil {
		return fmt.Errorf("output.Weather nil")
	}
	if len(output.Weather.Weather) == 0 {
		return fmt.Errorf("output.Weather.Weather size zero")
	}
	return nil
}
