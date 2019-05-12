package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func GetAPI() *Result {
	baseUrl := "http://api.openweathermap.org/data/2.5/forecast?"
	key := os.Getenv("API_KEY")
	values := url.Values{}

	values.Add("appid", key)      // OpenWeatherのAPIKey
	values.Add("lat", "36.5286")  // 緯度
	values.Add("lon", "136.6283") // 経度
	values.Add("units", "metric") // 温度（℃）

	response, err := http.Get(baseUrl + values.Encode())
	if err != nil {
		log.Fatalf("Connection Error: %v", err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Connection Error: %v", err)
	}

	jsonBytes := ([]byte)(body)
	date := new(Result)
	if err := json.Unmarshal(jsonBytes, date); err != nil {
		log.Fatalf("Connection Error: %v", err)
	}
	return date
}
