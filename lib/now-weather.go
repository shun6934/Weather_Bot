package lib

import (
	"math"
	"time"
)

func GetNowWeather() string {
	comment := ""
	date := GetAPI()

	weather, description := DiscriminateWeather(date)

	switch weather {
	case "Clear":
		comment = "現在、綺麗に晴れています。" + "(" + description + ")"
	case "Clouds":
		comment = "現在、曇りです。" + "(" + description + ")"
	case "Rain":
		comment = "現在、雨です。" + "(" + description + ")"
	case "Snow":
		comment = "現在、雪です。" + "(" + description + ")"
	default:
		comment = weather + "(" + description + ")"
	}

	return comment
}

func DiscriminateWeather(date *Result) (string, string) {
	weather := ""
	description := ""

	const timeLayout = "2006-01-02 15:04:05"
	now := time.Now()

	if date.List != nil {
		var prevTime time.Time
		var prevWeather string
		var prevDescription string

		for _, getTime := range date.List {
			loc, _ := time.LoadLocation("Asia/Tokyo")
			dateTime, _ := time.ParseInLocation(timeLayout, getTime.DtTxt, loc) // string -> time.Time

			beforeDuration := math.Abs(now.Sub(prevTime).Hours())
			afterDuration := math.Abs(now.Sub(dateTime).Hours())

			if now.Before(dateTime) {
				if beforeDuration < afterDuration {
					weather = prevWeather
					description = prevDescription
				} else {
					weather = getTime.Weather[0].Main
					description = getTime.Weather[0].Description
				}
				break
			}
			prevTime = dateTime
			prevWeather = getTime.Weather[0].Main
			prevDescription = getTime.Weather[0].Description
		}
	}
	return weather, description
}
