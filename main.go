package main

import (
	"line-bot-weather/frameworksanddrivers"
	"log"
	"net/http"
	"os"
)

func main() {
	factoryMethod := frameworksanddrivers.NewFactoryMethod()
	http.HandleFunc("/", factoryMethod.GetWeatherContller(os.Getenv("APPID")).GetWeather)
	http.HandleFunc("/callback", factoryMethod.GetWeatherLineContller(os.Getenv("APPID"), os.Getenv("CHANNEL_SECRET"), os.Getenv("CHANNEL_TOKEN")).GetWeatherLine)
	port := os.Getenv("PORT")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		log.Println("Listening...")
	}
}
