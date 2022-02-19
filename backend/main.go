package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	css := http.FileServer(http.Dir("../client/style"))
	http.Handle("/static/", http.StripPrefix("/static/", css))

	twitchClient := TwitchClient{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
		ClientID:     "m8uvbc0xacxewrobzsfa5ps6al2dlb",
		ClientSecret: "ou9zyyzat2c0vfpyc41wqax77rvzbm",
		Scopes:       []string{"user:read:follows"},
	}

	handler := Handler{Client: twitchClient.Client}
	handler.HandleTemplates("../templates")

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
		id := result.Id
		emotes, err := twitchClient.GetEmotes(id)
		if err != nil {
			log.Fatal(err)
		}
		handler.ExecuteTemplate(w, "index", ToJSON(*emotes))
	})
	err = handler.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
