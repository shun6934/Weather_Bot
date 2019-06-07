package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nlopes/slack"
	"github.com/shun6934/Weather_Bot/lib"
)

func main() {
	http.HandleFunc("/weather", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	s, err := slack.SlashCommandParse(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !s.ValidateToken(os.Getenv("VERIFICATION_TOKEN")) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	switch s.Command {
	case "/weather":
		weatherParams := &slack.Msg{Text: lib.GetNowWeather(), ResponseType: "in_channel"}
		tempertureParams := &slack.Msg{Text: lib.GetNowTemperture(), ResponseType: "in_channel"}

		w.Header().Set("Content-Type", "application/json")

		result := fmt.Sprintf("%v\n%v", weatherParams.Text, tempertureParams.Text)

		w.Write([]byte(result))
	default:
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
