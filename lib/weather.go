package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

func GetWeather() string {
	values := url.Values{}
	baseUrl := "http://api.openweathermap.org/data/2.5/forecast?"

	values.Add("appid", "[APIKey]") // OpenWeatherのAPIKey
	values.Add("lat", "36.5286")                            // 緯度
	values.Add("lon", "136.6283")                           // 経度

	weather := ParseJson(baseUrl + values.Encode())
	return weather
}

func ParseJson(url string) string {
	weather := ""

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
	}

	if data.List != nil {
		weather = data.List[0].Weather[0].Main
	}
	return weather
}
