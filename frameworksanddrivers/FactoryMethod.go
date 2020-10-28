package frameworksanddrivers

import (
	"line-bot-weather/interfaceadapters/controllers"
	"line-bot-weather/interfaceadapters/presenters"
	"line-bot-weather/usecases"
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
func (factoryMethod FactoryMethod) GetWeatherContller(appid string) controllers.WeatherContllerIF {
	return controllers.NewWeatherContller(usecases.NewWeatherInteractor(NewOpenWeatherMapRepository(), presenters.NewWeatherPresenter()), appid)
}

// GetWeatherLineContller 天気情報取得メソッド
func (factoryMethod FactoryMethod) GetWeatherLineContller(appid string, channelSecret string, channelToken string) controllers.WeatherContllerIF {
	return controllers.NewWeatherLineContller(usecases.NewWeatherInteractor(NewOpenWeatherMapRepository(), presenters.NewWeatherPresenter()), appid, channelSecret, channelToken)
}
