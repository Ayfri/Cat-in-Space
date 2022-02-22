package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type DataState struct {
	BestChannels []UserData
	DreamSmp     []UserData
	Results      []UserData
	Search       string
	Streamer     UserData
	ShowStreamer bool
}

func main() {
	css := http.FileServer(http.Dir("../client/style"))
	http.Handle("/static/", http.StripPrefix("/static/", css))

	js := http.FileServer(http.Dir("../client/scripts"))
	http.Handle("/js/", http.StripPrefix("/js/", js))

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	dataState := DataState{}

	twitchClient := TwitchClient{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"user:read:follows"},
	}

	handler := Handler{Client: twitchClient.Client}
	handler.HandleTemplates("../templates")

	err = twitchClient.FetchToken()
	if err != nil {
		log.Fatal(err)
	}

	DreamSmp := []string{"dreamwastaken", "georgenotfound", "sapnap", "badboyhalo", "tommyinnit", "tubbo", "ranboolive", "karljacobs", "nihachu", "quackity"}
	BestChannel := []string{"ayfri1015", "xhmyjae", "antaww", "kerrr_z", "amouranth"}

	for _, s := range DreamSmp {
		userdata, _ := twitchClient.GetUserByLogin(s)
		dataState.DreamSmp = append(dataState.DreamSmp, *userdata)
	}

	for _, s := range BestChannel {
		userdata, _ := twitchClient.GetUserByLogin(s)
		dataState.BestChannels = append(dataState.BestChannels, *userdata)
	}

	handler.HandleRoute("/", func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()
		if queries.Has("name") {
			streamer, _ := twitchClient.GetUserByLogin(queries.Get("name"))
			dataState.Streamer = *streamer
			if err != nil {
				log.Fatal(err)
			}
			dataState.ShowStreamer = true
		} else if r.Method == "GET" {
			dataState.Search = ""
			dataState.ShowStreamer = false
		}
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				log.Fatal(err)
			}
			dataState.ShowStreamer = false
			dataState.Search = r.Form.Get("search")
			if dataState.Search == "" {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}

			results, err := twitchClient.SearchChannelsAndFetch(dataState.Search)
			if err != nil {
				log.Fatal(err)
			}
			dataState.Results = *results
		}
		handler.ExecuteTemplate(w, "index", dataState)
	})

	err = handler.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
