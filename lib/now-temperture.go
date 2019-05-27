package lib

import (
	"strconv"
)

func GetNowTemperture() string {
	comment := ""
	temperture := ""

	date, _ := GetAPI("http://api.openweathermap.org/data/2.5/weather?")

	if date != nil {
		temperture = strconv.FormatFloat(date.Main.Temp, 'g', 4, 64)
		comment = "現在の気温は、" + temperture + "度です。"
	}
	return comment
}
