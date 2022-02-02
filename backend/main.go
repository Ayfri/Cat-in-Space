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
	err := twitchClient.FetchToken()
	if err != nil {
		log.Fatal(err)
	}
	var result *UserData
	result, err = twitchClient.GetUserByLogin("Ayfri1015")
	log.Printf(`
Name: %s
ID: %s
Type: %s
Views: %d
`, result.DisplayName, result.Id, result.BroadcasterType, result.ViewCount)
}
