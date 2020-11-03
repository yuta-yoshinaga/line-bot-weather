package frameworksanddrivers

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// TestGetCurrentWeather GetCurrentWeatherのテスト
func TestGetCurrentWeather(t *testing.T) {
	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		fmt.Println("env file not found.")
	}
	openWeatherMapRepository := NewOpenWeatherMapRepository()
	t.Run("success NewOpenWeatherMapRepository()", func(t *testing.T) {
		if openWeatherMapRepository == nil {
			t.Errorf("NG")
		}
	})
	t.Run("failed GetCurrentWeather()", func(t *testing.T) {
		query := ""
		currentWeather := openWeatherMapRepository.GetCurrentWeather(query)
		if currentWeather == nil {
			t.Errorf("NG")
			return
		}
		assert.Equal(t, 401, currentWeather.Cod)
	})
	t.Run("success GetCurrentWeather()", func(t *testing.T) {
		query := ""
		query += "q=tokyo,JP"
		query += "&appid=" + os.Getenv("APPID")
		query += "&lang=ja"
		query += "&units=metric"
		currentWeather := openWeatherMapRepository.GetCurrentWeather(query)
		if currentWeather == nil {
			t.Errorf("NG")
			return
		}
		assert.Equal(t, 200, currentWeather.Cod)
	})
}
