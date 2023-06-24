package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yuta-yoshinaga/line-bot-weather/frameworksanddrivers"
)

func main() {
	factoryMethod := frameworksanddrivers.NewFactoryMethod()
	http.HandleFunc("/", factoryMethod.GetWeatherContller(os.Getenv("APPID"), os.Getenv("WEATHER_URL")).GetWeather)
	http.HandleFunc("/callback", factoryMethod.GetWeatherLineContller(os.Getenv("APPID"), os.Getenv("WEATHER_URL"), os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_TOKEN")).GetWeatherLine)
	port := os.Getenv("PORT")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		log.Println("Listening...")
	}
}
