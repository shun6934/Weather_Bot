package lib

import (
	"math"
	"strconv"
	"time"
)

func GetNowTemperture() string {
	comment := ""
	temperture := ""

	date := GetAPI()

	const timeLayout = "2006-01-02 15:04:05"
	now := time.Now()

	if date.List != nil {
		var prevTime time.Time
		var prevTem float64

		for _, getTime := range date.List {
			loc, _ := time.LoadLocation("Asia/Tokyo")
			dateTime, _ := time.ParseInLocation(timeLayout, getTime.DtTxt, loc) // string -> time.Time

			beforeDuration := math.Abs(now.Sub(prevTime).Hours())
			afterDuration := math.Abs(now.Sub(dateTime).Hours())

			if now.Before(dateTime) {
				if beforeDuration < afterDuration {
					temperture = strconv.FormatFloat(prevTem, 'g', 4, 64)
				} else {
					temperture = strconv.FormatFloat(getTime.Main.Temp, 'g', 4, 64)
				}
				break
			}
			prevTime = dateTime
			prevTem = getTime.Main.Temp
		}
		comment = "現在の気温は、" + temperture + "度です。"
	}
	return comment
}
