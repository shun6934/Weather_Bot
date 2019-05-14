package lib

import (
	"strconv"
	"time"
)

func GetNowTemperture() string {
	comment := ""
	var temperture string

	date := GetAPI()

	const timeLayout = "2006-01-02 15:04:05"
	now := time.Now()

	if date.List != nil {
		for _, getTime := range date.List {
			loc, _ := time.LoadLocation("Asia/Tokyo")
			t, _ := time.ParseInLocation(timeLayout, getTime.DtTxt, loc) // string -> time.Time
			if !now.After(t) {
				temperture = strconv.FormatFloat(getTime.Main.Temp, 'g', 4, 64)
				break
			}
		}
		comment = "現在の気温は、" + temperture + "度です。"
	}
	return comment
}
