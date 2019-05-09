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

type Result struct {
	List []struct {
		Dt   int `json:"dt"`
		Main struct {
			Temp     float64 `json:"temp"`
			TempMin  float64 `json:"temp_min"`
			TempMax  float64 `json:"temp_max"`
			Humidity int     `json:"humidity"`
		} `json:"main"`
		Weather []struct {
			Main        string `json:"main"`
			Description string `json:"description"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   float64 `json:"deg"`
		} `json:"wind"`
		DtTxt string `json:"dt_txt"`
	} `json:"list"`
}

func GetWeather(now string) string {
	comment := ""
	key := os.Getenv("API_KEY")

	values := url.Values{}
	baseUrl := "http://api.openweathermap.org/data/2.5/forecast?"

	values.Add("appid", key)      // OpenWeatherのAPIKey
	values.Add("lat", "36.5286")  // 緯度
	values.Add("lon", "136.6283") // 経度

	weather := ParseJson(baseUrl+values.Encode(), now)

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
	return comment
}

func ParseJson(url string, now string) string {
	weather := ""
	const timeLayout = "2006-01-02 15:00:00"
	nowToTime, _ := time.Parse(timeLayout, now) // string -> time.Time

	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Connection Error: %v", err)
		return "取得できませんでした"
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Connection Error: %v", err)
		return "取得できませんでした"
	}

	jsonBytes := ([]byte)(body)
	data := new(Result)
	if err := json.Unmarshal(jsonBytes, data); err != nil {
		log.Fatalf("Connection Error: %v", err)
		return "取得できませんでした"
	}

	if data.List != nil {
		for _, getTime := range data.List {
			t, _ := time.Parse(timeLayout, getTime.DtTxt) // string -> time.Time
			if !nowToTime.After(t) {
				weather = getTime.Weather[0].Main
				break
			}
		}
	}
	return weather
}
