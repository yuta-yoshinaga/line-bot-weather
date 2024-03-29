package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/yuta-yoshinaga/line-bot-weather/usecases"

	"github.com/line/line-bot-sdk-go/linebot"
)

// WeatherLineContller LINE天気情報取得コントローラ
type WeatherLineContller struct {
	WeatherContller
	ChannelSecret string
	ChannelToken  string
}

// NewWeatherLineContller コンストラクタ
func NewWeatherLineContller(usecase usecases.WeatherInteractorIF, Appid string, ChannelSecret string, ChannelToken string) *WeatherLineContller {
	weatherContller := new(WeatherLineContller)
	weatherContller.usecase = usecase
	weatherContller.Appid = Appid
	weatherContller.ChannelSecret = ChannelSecret
	weatherContller.ChannelToken = ChannelToken
	return weatherContller
}

// GetWeatherLine 天気情報取得メソッド
func (weatherContller WeatherLineContller) GetWeatherLine(w http.ResponseWriter, r *http.Request) {
	bot, err := linebot.New(
		weatherContller.ChannelSecret,
		weatherContller.ChannelToken,
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
	input := new(usecases.WeatherInput)
	input.Appid = weatherContller.Appid
	for _, event := range events {
		fmt.Println(event)
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				fmt.Println(message)
				if event.ReplyToken == "00000000000000000000000000000000" {
					return
				}
				input.City = getCityNameASCII(message.Text)
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(weatherContller.usecase.GetCurrentWeather(input))).Do(); err != nil {
					log.Print(err)
				}
			case *linebot.LocationMessage:
				fmt.Println(message)
				if event.ReplyToken == "00000000000000000000000000000000" {
					return
				}
				input.Lat = strconv.FormatFloat(message.Latitude, 'f', -1, 64)
				input.Lon = strconv.FormatFloat(message.Longitude, 'f', -1, 64)
				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(weatherContller.usecase.GetCurrentWeather(input))).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}

func getCityNameASCII(text string) string {
	res := ""
	if strings.HasPrefix(text, "北海道") {
		res = "hokkaido,JP"
	} else if strings.HasPrefix(text, "青森") {
		res = "aomori,JP"
	} else if strings.HasPrefix(text, "岩手") {
		res = "iwate,JP"
	} else if strings.HasPrefix(text, "宮城") {
		res = "miyagi,JP"
	} else if strings.HasPrefix(text, "秋田") {
		res = "akita,JP"
	} else if strings.HasPrefix(text, "山形") {
		res = "yamagata,JP"
	} else if strings.HasPrefix(text, "福島") {
		res = "fukushima,JP"
	} else if strings.HasPrefix(text, "茨城") {
		res = "ibaraki,JP"
	} else if strings.HasPrefix(text, "栃木") {
		res = "tochigi,JP"
	} else if strings.HasPrefix(text, "群馬") {
		res = "gunma,JP"
	} else if strings.HasPrefix(text, "埼玉") {
		res = "saitama,JP"
	} else if strings.HasPrefix(text, "千葉") {
		res = "chiba,JP"
	} else if strings.HasPrefix(text, "東京") {
		res = "tokyo,JP"
	} else if strings.HasPrefix(text, "神奈川") {
		res = "kanagawa,JP"
	} else if strings.HasPrefix(text, "新潟") {
		res = "niigata,JP"
	} else if strings.HasPrefix(text, "富山") {
		res = "toyama,JP"
	} else if strings.HasPrefix(text, "石川") {
		res = "ishikawa,JP"
	} else if strings.HasPrefix(text, "福井") {
		res = "fukui,JP"
	} else if strings.HasPrefix(text, "山梨") {
		res = "yamanashi,JP"
	} else if strings.HasPrefix(text, "長野") {
		res = "nagano,JP"
	} else if strings.HasPrefix(text, "岐阜") {
		res = "gifu,JP"
	} else if strings.HasPrefix(text, "静岡") {
		res = "shizuoka,JP"
	} else if strings.HasPrefix(text, "愛知") {
		res = "aichi,JP"
	} else if strings.HasPrefix(text, "三重") {
		res = "mie,JP"
	} else if strings.HasPrefix(text, "滋賀") {
		res = "shiga,JP"
	} else if strings.HasPrefix(text, "京都") {
		res = "kyoto,JP"
	} else if strings.HasPrefix(text, "大阪") {
		res = "osaka,JP"
	} else if strings.HasPrefix(text, "兵庫") {
		res = "hyogo,JP"
	} else if strings.HasPrefix(text, "奈良") {
		res = "nara,JP"
	} else if strings.HasPrefix(text, "和歌山") {
		res = "wakayama,JP"
	} else if strings.HasPrefix(text, "鳥取") {
		res = "tottori,JP"
	} else if strings.HasPrefix(text, "島根") {
		res = "shimane,JP"
	} else if strings.HasPrefix(text, "岡山") {
		res = "okayama,JP"
	} else if strings.HasPrefix(text, "広島") {
		res = "hiroshima,JP"
	} else if strings.HasPrefix(text, "山口") {
		res = "yamaguchi,JP"
	} else if strings.HasPrefix(text, "徳島") {
		res = "tokushima,JP"
	} else if strings.HasPrefix(text, "香川") {
		res = "kagawa,JP"
	} else if strings.HasPrefix(text, "愛媛") {
		res = "ehime,JP"
	} else if strings.HasPrefix(text, "高知") {
		res = "kochi,JP"
	} else if strings.HasPrefix(text, "福岡") {
		res = "fukuoka,JP"
	} else if strings.HasPrefix(text, "佐賀") {
		res = "saga,JP"
	} else if strings.HasPrefix(text, "長崎") {
		res = "nagasaki,JP"
	} else if strings.HasPrefix(text, "熊本") {
		res = "kumamoto,JP"
	} else if strings.HasPrefix(text, "大分") {
		res = "oita,JP"
	} else if strings.HasPrefix(text, "宮崎") {
		res = "miyazaki,JP"
	} else if strings.HasPrefix(text, "鹿児島") {
		res = "kagoshima,JP"
	} else if strings.HasPrefix(text, "沖縄") {
		res = "okinawa,JP"
	} else {
		res = "tokyo,JP"
	}
	return res
}
