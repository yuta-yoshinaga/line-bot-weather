package main

import (
	"fmt"
	"line-bot-weather/weather"
	"log"
	"net/http"
	"os"
)

func getWeather(w http.ResponseWriter, r *http.Request) {
	appid := os.Getenv("APPID")
	query := "q=Tokyo,jp"
	query += "&appid=" + appid
	query += "&lang=ja"
	query += "&units=metric"
	data := weather.NewCurrentWeather(query)
	fmt.Println(data)
	fmt.Fprintf(w, data.GetWeatherText())
}

func main() {
	http.HandleFunc("/", getWeather)
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		log.Println("Listening...")
	}
}
