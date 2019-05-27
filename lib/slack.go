package lib

import (
	"log"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

func Run(api *slack.Client) int {
	var nowWeather string
	var nowTemperture string

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
					nowWeather = GetNowWeather()
					nowTemperture = GetNowTemperture()

					rtm.SendMessage(rtm.NewOutgoingMessage(nowWeather, ev.Channel))
					rtm.SendMessage(rtm.NewOutgoingMessage(nowTemperture, ev.Channel))
				}
			case *slack.InvalidAuthEvent:
				log.Println("Invalid credentials")
				return 1
			}
		}
	}
}
