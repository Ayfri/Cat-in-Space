package main

import (
	"log"
	"net/http"
	"os"
	"sort"
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

	handler.HandleResourcesDir("../client/style", "/static/")
	handler.HandleResourcesDir("../client/scripts", "/js/")
	handler.HandleResourcesDir("../backend/resources", "/resources/")

	err = twitchClient.FetchToken()
	if err != nil {
		log.Fatal(err)
	}

	DreamSmp := []string{"Dream", "georgenotfound", "sapnap", "badboyhalo", "tommyinnit", "tubbo", "ranboolive", "karljacobs", "nihachu", "quackity", "antfrost"}
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
		if r.URL.Path != "/" || r.URL.String() == "/?" {
			url := "/"
			query := r.URL.RawQuery
			if query != "" {
				url += "?" + query
			}
			http.Redirect(w, r, url, http.StatusFound)
			return
		}
		queries := r.URL.Query()

		if len(queries) == 0 {
			dataState.Search = ""
			dataState.ShowStreamer = false
		}

		if queries.Has("query") {
			search := queries.Get("query")
			if dataState.Search != search {
				twitchClient.Cursor = ""
				dataState.Search = search
			}

			results, err := twitchClient.SearchChannelsAndFetch(dataState.Search, twitchClient.Cursor)
			if err != nil {
				log.Fatal(err)
			}

			if dataState.Search != search {
				dataState.Results = *results
			} else {
				for _, data := range *results {
					resultsTest := false
					for _, dataa := range dataState.Results {
						if data.Id == dataa.Id {
							resultsTest	= true
						}
					}
					if !resultsTest {
						dataState.Results = append(dataState.Results, data)
					}
				}
			}
			
			sort.Slice(dataState.Results, func(i, j int) bool {
				return dataState.Results[i].ViewCount > dataState.Results[j].ViewCount
			})
		}

		if queries.Has("name") {
			streamer, err := twitchClient.GetUserByLogin(queries.Get("name"))
			if err != nil {
				log.Fatal(err)
			}
			if streamer != nil {
				streamer.GetEmotes(twitchClient)
				dataState.Streamer = *streamer
			}
			dataState.ShowStreamer = true
		}

		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				log.Fatal(err)
			}

			if r.Form.Has("load-more") {
				results, err := twitchClient.SearchChannelsAndFetch(dataState.Search, twitchClient.Cursor)
				if err != nil {
					log.Fatal(err)
				}

				for _, s := range *results {
					dataState.Results = append(dataState.Results, s)
				}

				sort.Slice(*results, func(i, j int) bool {
					return (*results)[i].ViewCount > (*results)[j].ViewCount
				})
			} else {
				twitchClient.Cursor = ""
				dataState.ShowStreamer = false
				dataState.Results = []UserData{}
				dataState.Search = r.Form.Get("search")
				if dataState.Search == "" {
					http.Redirect(w, r, "/", http.StatusSeeOther)
				} else {
					http.Redirect(w, r, "/?query="+dataState.Search, http.StatusSeeOther)
				}
				return
			}
		}

		handler.ExecuteTemplate(w, "index", dataState)
	})

	err = handler.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
