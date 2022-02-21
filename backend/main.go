package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type DataState struct {
	Search  string
	Results []UserData
}

func main() {
	css := http.FileServer(http.Dir("../client/style"))
	http.Handle("/static/", http.StripPrefix("/static/", css))

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

	handler.HandleRoute("/", func(w http.ResponseWriter, r *http.Request) {
		search := r.FormValue("search")
		if search != "" {
			dataState.Search = search
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
