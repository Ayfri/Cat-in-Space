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

	handler := Handler{
		Client: twitchClient.Client,
	}
	var result *UserData
	result, err = twitchClient.GetUserByLogin("Ayfri1015")
	handler.HandleTemplates("./templates")
	println(handler.tree.DefinedTemplates())
	for name, template := range handler.Templates {
		log.Println(name, template)
	}
	handler.HandleRoute("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request:", r.URL.Path)
		handler.ExecuteTemplate(w, "index", ToJSON(*result))
	})
	err = handler.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
