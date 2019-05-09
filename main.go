package main

import (
	"fmt"
	"time"
	"github.com/shun6934/Weather_Bot/lib"
)

func main() {
	getTime := time.Now()
	const timeLayout = "2006-01-02 15:00:00"
	now := getTime.Format(timeLayout) // tiem.Time -> string

	result := lib.GetWeather(now)
	fmt.Printf(result)
}