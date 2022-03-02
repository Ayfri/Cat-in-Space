package main

import (
	"net/http"
	"strconv"
	"strings"
)

type TwitchClient struct {
	Client       *http.Client
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Scopes       []string
	Token        *TokenResponse
	Cursor string
}

func (client *TwitchClient) FetchToken() error {
	requester := Requester{
		Client: *client.Client,
		Method: "POST",
		URL:    "https://id.twitch.tv/oauth2/token",
		URLParams: map[string]string{
			"client_id":     client.ClientID,
			"client_secret": client.ClientSecret,
			"grant_type":    "client_credentials",
			"scope":         strings.Join(client.Scopes, " "),
		},
	}
	result := &TokenResponse{}
	err := requester.DoRequestTo(result)
	if err != nil {
		return err
	}
	client.Token = result
	return nil
}

func (client *TwitchClient) GetEmotes(id string) (*EmoteResponse, error) {
	requester := Requester{
		Client: *client.Client,
		Headers: map[string]string{
			"Authorization": client.Token.GetFormattedToken(),
			"Client-ID":     client.ClientID,
		},
		Method: "GET",
		URL:    "https://api.twitch.tv/helix/chat/emotes",
		URLParams: map[string]string{
			"broadcaster_id": id,
		},
	}
	result := &EmoteResponse{}
	err := requester.DoRequestTo(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *TwitchClient) GetFollowers(id string, count int) (*FollowersResponse, error) {
	requester := Requester{
		Client: *client.Client,
		Headers: map[string]string{
			"Authorization": client.Token.GetFormattedToken(),
			"Client-ID":     client.ClientID,
		},
		Method: "GET",
		URL:    "https://api.twitch.tv/helix/users/follows",
		URLParams: map[string]string{
			"first": strconv.Itoa(count),
			"to_id": id,
		},
	}
	result := &FollowersResponse{}
	err := requester.DoRequestTo(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (client *TwitchClient) GetUserByLogin(login string) (*UserData, error) {
	requester := Requester{
		Client: *client.Client,
		Headers: map[string]string{
			"Authorization": client.Token.GetFormattedToken(),
			"Client-ID":     client.ClientID,
		},
		Method: "GET",
		URL:    "https://api.twitch.tv/helix/users",
		URLParams: map[string]string{
			"login": login,
		},
	}
	result := &UserDataResponse{}
	err := requester.DoRequestTo(result)
	if err != nil {
		return nil, err
	}
	if len(result.Data) == 0 {
		return nil, nil
	}
	return &result.Data[0], nil
}

func (client *TwitchClient) GetUsersById(ids []string) (*[]UserData, error) {
	requester := Requester{
		Client: *client.Client,
		Headers: map[string]string{
			"Authorization": client.Token.GetFormattedToken(),
			"Client-ID":     client.ClientID,
		},
		Method: "GET",
		URL:    "https://api.twitch.tv/helix/users",
	}
	for _, login := range ids {
		requester.URLParamsArray = append(requester.URLParamsArray, Pair{"id", login})
	}

	result := &UserDataResponse{}
	err := requester.DoRequestTo(result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}

func (client *TwitchClient) GetUsers(users *[]UserData) (*[]UserData, error) {
	var names []string
	for _, user := range *users {
		names = append(names, user.Id)
	}
	return client.GetUsersById(names)
}

func (client *TwitchClient) SearchChannels(query string, after string) (*[]UserData, error) {
	requester := Requester{
		Client: *client.Client,
		Headers: map[string]string{
			"Authorization": client.Token.GetFormattedToken(),
			"Client-ID":     client.ClientID,
		},
		Method: "GET",
		URL:    "https://api.twitch.tv/helix/search/channels",
		URLParams: map[string]string{
			"query": query,
			"first": "99",
			"after": after,
		},
	}
	result := &UserDataResponse{}
	err := requester.DoRequestTo(result)
	if err != nil {
		return nil, err
	}
	client.Cursor = result.Pagination.Cursor
	return &result.Data, nil
}

func (client *TwitchClient) SearchChannelsAndFetch(query string, after string) (*[]UserData, error) {
	channels, err := client.SearchChannels(query, after)
	if err != nil {
		return nil, err
	}
	if len(*channels) == 0 {
		return channels, nil
	}
	return client.GetUsers(channels)
}
