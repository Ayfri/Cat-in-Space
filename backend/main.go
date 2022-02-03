package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	twitchClient := TwitchClient{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
		ClientID:     "My little id :>",
		ClientSecret: "Ho boy we should use .env or something before someone accidentally pushes its token",
		Scopes:       []string{"user:read:follows"},
	}

	handler := Handler{Client: twitchClient.Client}
	handler.HandleTemplates("./templates")

	err := twitchClient.FetchToken()
	if err != nil {
		log.Fatal(err)
	}
	handler.HandleRoute("/", func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()
		user := queries.Get("user")
		if user == "" {
			user = "Ayfri1015"
		}
		log.Println("User:", user)
		result, err := twitchClient.GetUserByLogin(user)
		if err != nil {
			log.Fatal(err)
		}
		handler.ExecuteTemplate(w, "index", ToJSON(*result))
	})
	err = handler.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
