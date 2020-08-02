package main

import (
	"fmt"
	"line-bot-weather/weather"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/line/line-bot-sdk-go/linebot"
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

func getWeatherLine(w http.ResponseWriter, r *http.Request) {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	events, err := bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}
	for _, event := range events {
		fmt.Println(event)
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if event.ReplyToken == "00000000000000000000000000000000" {
					return
				}
				appid := os.Getenv("APPID")
				query := "q=Tokyo,jp"
				query += "&appid=" + appid
				query += "&lang=ja"
				query += "&units=metric"
				data := weather.NewCurrentWeather(query)
				fmt.Println(message)
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(data.GetWeatherText())).Do(); err != nil {
					log.Print(err)
				}
			case *linebot.LocationMessage:
				if event.ReplyToken == "00000000000000000000000000000000" {
					return
				}
				appid := os.Getenv("APPID")
				query := "lat=" + strconv.FormatFloat(message.Latitude, 'f', -1, 64)
				query += "&lon=" + strconv.FormatFloat(message.Longitude, 'f', -1, 64)
				query += "&appid=" + appid
				query += "&lang=ja"
				query += "&units=metric"
				data := weather.NewCurrentWeather(query)
				fmt.Println(message)
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(data.GetWeatherText())).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}

func main() {
	http.HandleFunc("/", getWeather)
	http.HandleFunc("/callback", getWeatherLine)
	port := os.Getenv("PORT")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		log.Println("Listening...")
	}
}
