package frameworksanddrivers

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// TestGetCurrentWeather1 GetCurrentWeatherのテスト
func TestOpenWeatherMapRepository_GetCurrentWeather(t *testing.T) {
	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		fmt.Println("env file not found.")
	}
	openWeatherMapRepository := NewOpenWeatherMapRepository(os.Getenv("WEATHER_URL"))
	t.Run("success NewOpenWeatherMapRepository()", func(t *testing.T) {
		if openWeatherMapRepository == nil {
			t.Errorf("NG")
		}
	})
	t.Run("success GetCurrentWeather()", func(t *testing.T) {
		query := ""
		query += "q=tokyo,JP"
		query += "&appid=" + os.Getenv("APPID")
		query += "&lang=ja"
		query += "&units=metric"
		currentWeather := openWeatherMapRepository.GetCurrentWeather(query)
		assert.NotEqual(t, nil, currentWeather)
		assert.Equal(t, 200, currentWeather.Cod)
	})
	t.Run("failed GetCurrentWeather()", func(t *testing.T) {
		query := ""
		currentWeather := openWeatherMapRepository.GetCurrentWeather(query)
		assert.NotEqual(t, nil, currentWeather)
		assert.Equal(t, 401, currentWeather.Cod)
	})
	t.Run("failed GetCurrentWeather()", func(t *testing.T) {
		openWeatherMapRepository := NewOpenWeatherMapRepository("")
		query := ""
		currentWeather := openWeatherMapRepository.GetCurrentWeather(query)
		assert.Empty(t, currentWeather)
	})
}
