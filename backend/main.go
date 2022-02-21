package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type DataState struct {
	DreamSmp []UserData
	BestChannels []UserData
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

	dataState := DataState{}

	for _, s := range DreamSmp {
		userdata, _ := twitchClient.GetUserByLogin(s)
		dataState.DreamSmp = append(dataState.DreamSmp, *userdata)
	}

	for _, s := range BestChannel {
		userdata, _ := twitchClient.GetUserByLogin(s)
		dataState.BestChannels = append(dataState.BestChannels, *userdata)
	}

	handler.HandleRoute("/", func(w http.ResponseWriter, r *http.Request) {
		//queries := r.URL.Query()
		//user := queries.Get("user")
		//if user == "" {
		//	user = "Ayfri1015"
		//}
		//log.Println("User:", user)
		//result, err := twitchClient.GetUserByLogin(user)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//id := result.Id
		//emotes, err := twitchClient.GetEmotes(id)
		//if err != nil {
		//	log.Fatal(err)
		//}
		handler.ExecuteTemplate(w, "index", dataState)
	})

	err = handler.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
