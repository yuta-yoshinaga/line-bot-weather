package usecases

// WeatherInteractor
type WeatherInteractor struct {
	repository OpenWeatherMapRepositoryIF
	presenter  WeatherOutputIF
}

// NewWeatherInteractor コンストラクタ
func NewWeatherInteractor(repository OpenWeatherMapRepositoryIF, presenter WeatherOutputIF) *WeatherInteractor {
	weatherInteractor := new(WeatherInteractor)
	weatherInteractor.repository = repository
	weatherInteractor.presenter = presenter
	return weatherInteractor
}

// GetCurrentWeather 天気情報取得メソッド
func (weatherInteractor WeatherInteractor) GetCurrentWeather(input *WeatherInput) string {
	weatherOutput := new(WeatherOutput)
	query := "q=Tokyo,jp"
	query += "&appid=" + input.Appid
	query += "&lang=ja"
	query += "&units=metric"
	weatherOutput.Weather = weatherInteractor.repository.GetCurrentWeather(query)
	return weatherInteractor.presenter.GetWeatherText(weatherOutput)
}
