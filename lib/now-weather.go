package lib

func GetNowWeather() string {
	comment := ""
	date, _ := GetAPI("http://api.openweathermap.org/data/2.5/weather?")

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

func DiscriminateWeather(date *WeatherResult) (string, string) {
	weather := ""
	description := ""

	if date != nil {
		weather = date.Weather[0].Main
		description = date.Weather[0].Description
	}
	return weather, description
}
