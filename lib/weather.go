package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

func GetWeather() string {
	comment := ""

	baseUrl := "http://api.openweathermap.org/data/2.5/forecast?"
	key := os.Getenv("API_KEY")
	values := url.Values{}

	values.Add("appid", key)      // OpenWeatherのAPIKey
	values.Add("lat", "36.5286")  // 緯度
	values.Add("lon", "136.6283") // 経度
	values.Add("units", "metric") // 温度（℃）

	weather, description := ParseJson(baseUrl + values.Encode())

	switch weather {
	case "Clear":
		comment = "綺麗に晴れています。"
	case "Clouds":
		comment = "曇りです。"
	case "Rain":
		comment = "雨です。"
	case "Snow":
		comment = "雪です。"
	default:
		comment = weather
	}

	return comment + "(" + description + ")"
}

func ParseJson(url string) (string, string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Connection Error: %v", err)
		return "データ無し", "取得できませんでした"
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Connection Error: %v", err)
		return "データ無し", "取得できませんでした"
	}

	jsonBytes := ([]byte)(body)
	data := new(Result)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		log.Fatalf("Connection Error: %v", err)
		return "データ無し", "取得できませんでした"
	}

	weather, description := DiscriminateWeather(data)

	return weather, description
}

func DiscriminateWeather(data *Result) (string, string) {
	weather := ""
	description := ""

	const timeLayout = "2006-01-02 15:04:05"
	now := time.Now()

	if data.List != nil {
		for _, getTime := range data.List {
			loc, _ := time.LoadLocation("Asia/Tokyo")
			t, _ := time.ParseInLocation(timeLayout, getTime.DtTxt, loc) // string -> time.Time
			if !now.After(t) {
				weather = getTime.Weather[0].Main
				description = getTime.Weather[0].Description
				break
			}
		}
	}
	return weather, description
}
