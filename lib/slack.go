package lib

import (
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/nlopes/slack"
)

func Run(api *slack.Client) int {
	getTime := time.Now()
	const timeLayout = "2006-01-02 15:00:00"
	now := getTime.Format(timeLayout) // tiem.Time -> string

	var result string

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
					result = rep.ReplaceAllString(ev.Text, "")
					result = GetWeather(now)
					rtm.SendMessage(rtm.NewOutgoingMessage(result, ev.Channel))
				}
			case *slack.InvalidAuthEvent:
				log.Println("Invalid credentials")
				return 1
			}
		}
	}
}
