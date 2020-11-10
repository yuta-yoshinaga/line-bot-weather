package frameworksanddrivers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"line-bot-weather/entities"
	"net/http"
)

// OpenWeatherMapRepository 天気情報取得リポジトリ
type OpenWeatherMapRepository struct {
	weatherURL string
}

// NewOpenWeatherMapRepository コンストラクタ
func NewOpenWeatherMapRepository(weatherURL string) *OpenWeatherMapRepository {
	openWeatherMapRepository := new(OpenWeatherMapRepository)
	openWeatherMapRepository.weatherURL = weatherURL
	return openWeatherMapRepository
}

// GetCurrentWeather 天気情報取得メソッド
func (openWeatherMapRepository OpenWeatherMapRepository) GetCurrentWeather(query string) *entities.CurrentWeather {
	currentWeather := new(entities.CurrentWeather)
	url := openWeatherMapRepository.weatherURL
	url += "?" + query
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP GET error:", err)
		return nil
	}
	defer resp.Body.Close()
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
	jsonBytes := ([]byte)(byteArray)
	if err := json.Unmarshal(jsonBytes, currentWeather); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return nil
	}
	return currentWeather
}
