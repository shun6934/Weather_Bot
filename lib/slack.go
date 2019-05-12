package lib

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/nlopes/slack"
)

func Run(api *slack.Client) int {
	var weather string

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Println("====Start====")
			case *slack.MessageEvent:
				if strings.Contains(ev.Text, os.Getenv("BOT_ID")) { // botのuserID = <@UJ79VEWF4>
					rep := regexp.MustCompile(os.Getenv(`BOT_ID`))
					weather = rep.ReplaceAllString(ev.Text, "")
					weather = GetWeather()
					rtm.SendMessage(rtm.NewOutgoingMessage(weather, ev.Channel))
				}
			case *slack.InvalidAuthEvent:
				log.Println("Invalid credentials")
				return 1
			}
		}
	}
}
