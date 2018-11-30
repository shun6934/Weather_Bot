package main

import (
	"github.com/nlopes/slack"
	"github.com/shun6934/slackbot/lib"
	"os"
)

func main() {
	token := slack.New(os.Getenv("SLACK_API?TOKEN"))
	os.Exit()
}