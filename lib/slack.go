package lib

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"net/url"
)

type Result struct {
	Feature []struct {
		ID       string `json:"Id"`
		Name     string `json:"Name"`
		Geometry struct {
			Type        string `json:"Type"`
			Coordinates string `json:"Coordinates"`
		} `json:"Geometry"`
		Property struct {
			WeatherList struct {
				Weather []struct {
					Type string `json:"Type"`
					Date string `json:"Date"`
				} `json:"Weather"`
			} `json:"WeatherList"`
		} `json:"Property"`
	} `json:"Feature"`
}

func main() {

	values := url.Values{}
	base_url := "https://map.yahooapis.jp/weather/V1/place?"

	values.Add("appid", "[APP-ID]")
	values.Add("coordinates", "35.663613,139.73229")
	values.Add("output", "json")

	resp, err := http.Get(base_url + values.Encode())
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonStr := ([]byte)(body)
	data := new(Result)
	if err := json.Unmarshal(jsonStr, data); err != nil {
		fmt.Println("JSON Unmarshal error:", err)
		return
	}
}
