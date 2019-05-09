package main

import (
	"os"

	"github.com/nlopes/slack"
	"github.com/shun6934/Weather_Bot/lib"
)

func main() {
	token := slack.New(os.Getenv("SLACK_API_TOKEN"))
	os.Exit(lib.Run(token))
}
