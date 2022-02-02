package main

import (
	"net/http"
	"strings"
	"time"
)

type TwitchClient struct {
	Client       *http.Client
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Scopes       []string
	Token        *TokenResponse
}

type TokenResponse struct {
	AccessToken string   `json:"access_token"`
	ExpiresIn   int      `json:"expires_in"`
	Scope       []string `json:"scope"`
	TokenType   string   `json:"token_type"`
}

func (response *TokenResponse) IsExpired() bool {
	return time.Now().Unix() > int64(response.ExpiresIn)
}

func (response *TokenResponse) GetExpiration() time.Time {
	return time.Now().Add(time.Duration(response.ExpiresIn) * time.Second)
}
func (response *TokenResponse) GetFormattedToken() string {
	t := response.TokenType
	t = strings.ToUpper(t[:1]) + t[1:]
	return t + " " + response.AccessToken
}

type UserData struct {
	Id              string    `json:"id"`
	Login           string    `json:"login"`
	DisplayName     string    `json:"display_name"`
	Type            string    `json:"type"`
	BroadcasterType string    `json:"broadcaster_type"`
	Description     string    `json:"description"`
	ProfileImageUrl string    `json:"profile_image_url"`
	OfflineImageUrl string    `json:"offline_image_url"`
	ViewCount       int       `json:"view_count"`
	Email           string    `json:"email"`
	CreatedAt       time.Time `json:"created_at"`
}

type UserDataResponse struct {
	Data []UserData `json:"data"`
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
	return &result.Data[0], nil
}
