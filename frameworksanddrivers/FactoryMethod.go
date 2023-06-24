package frameworksanddrivers

import (
	"github.com/yuta-yoshinaga/line-bot-weather/interfaceadapters/controllers"
	"github.com/yuta-yoshinaga/line-bot-weather/interfaceadapters/presenters"
	"github.com/yuta-yoshinaga/line-bot-weather/usecases"
)

// FactoryMethod ファクトリーメソッド
type FactoryMethod struct {
}

// NewFactoryMethod コンストラクタ
func NewFactoryMethod() *FactoryMethod {
	factoryMethod := new(FactoryMethod)
	return factoryMethod
}

// GetWeatherContller 天気情報取得メソッド
func (factoryMethod FactoryMethod) GetWeatherContller(appid string, url string) controllers.WeatherContllerIF {
	return controllers.NewWeatherContller(usecases.NewWeatherInteractor(NewOpenWeatherMapRepository(url),
		presenters.NewWeatherPresenter()),
		appid)
}

// GetWeatherLineContller 天気情報取得メソッド
func (factoryMethod FactoryMethod) GetWeatherLineContller(appid string, url string, channelSecret string, channelToken string) controllers.WeatherLineContllerIF {
	return controllers.NewWeatherLineContller(usecases.NewWeatherInteractor(NewOpenWeatherMapRepository(url),
		presenters.NewWeatherPresenter()),
		appid,
		channelSecret,
		channelToken)
}
