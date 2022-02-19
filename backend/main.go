package main

import (
	"log"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
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
		queries := r.URL.Query()
		user := queries.Get("user")
		if user == "" {
			user = "Ayfri1015"
		}
		// Search channels by user argument in URL
		result, err := twitchClient.SearchChannel(user)
		if err != nil {
			log.Fatal(err)
		}

		// Fetch all channels to get other data than Id & DisplayName
		result, err = twitchClient.GetUsers(result)
		if err != nil {
			log.Fatal(err)
		}

		// Sort channels by ViewCount
		sort.Slice(*result, func(i, j int) bool {
			return (*result)[i].ViewCount > (*result)[j].ViewCount
		})
		handler.ExecuteTemplate(w, "index", ToJSON(*result))
	})
	err = handler.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
