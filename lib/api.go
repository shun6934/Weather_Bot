package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func GetAPI(baseUrl string) (*WeatherResult, *ForecastResult) {
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
	weatherDate := new(WeatherResult)
	forecastDate := new(ForecastResult)
	if err := json.Unmarshal(jsonBytes, forecastDate); err != nil {
		log.Fatalf("Connection Error: %v", err)
	}
	if err := json.Unmarshal(jsonBytes, weatherDate); err != nil {
		log.Fatalf("Connection Error: %v", err)
	}
	return weatherDate, forecastDate
}
