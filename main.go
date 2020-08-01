package main

import (
	"fmt"
	"line-bot-weather/weather"
	"log"
	"net/http"
	"os"

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

func main() {
	http.HandleFunc("/", getWeather)
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
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
				}
			}
		}
	})
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {
		log.Println("Listening...")
	}
}
