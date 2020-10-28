package usecases

// WeatherInteractor 天気情報取得ユースケースインタラクター
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
	query := ""
	if input.City != "" {
		query += "q=" + input.City
	}
	if input.Lat != "" {
		query += "&lat=" + input.Lat
	}
	if input.Lon != "" {
		query += "&lon=" + input.Lon
	}
	query += "&appid=" + input.Appid
	query += "&lang=ja"
	query += "&units=metric"
	weatherOutput.Weather = weatherInteractor.repository.GetCurrentWeather(query)
	return weatherInteractor.presenter.GetWeatherText(weatherOutput)
}
