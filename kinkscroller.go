package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/caseymrm/menuet"
)

type nowPlaying struct {
	Playing string `json:"playing"`
}

func kinkScroller() {
	client := &http.Client{Timeout: 10 * time.Second}
	for {
		resp, err := client.Get("https://api.kink.nl/static/now-playing.json")
		if err != nil {
			panic(err)
		}

		np := nowPlaying{}
		err = json.NewDecoder(resp.Body).Decode(&np)
		resp.Body.Close()
		if err != nil {
			panic(err)
		}

		menuet.App().SetMenuState(&menuet.MenuState{
			Title: "ꓘ | " + np.Playing + " | K",
		})

		time.Sleep(20 * time.Second)
	}
}

func main() {
	go kinkScroller()
	menuet.App().RunApplication()
}
