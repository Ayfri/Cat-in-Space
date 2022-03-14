package main

import (
	"log"
	"strings"
	"time"
)

type EmoteResponse struct {
	Data     []EmoteData `json:"data"`
	Template string      `json:"template"`
}

type EmoteData struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Images struct {
		Url1X string `json:"url_1x"`
		Url2X string `json:"url_2x"`
		Url4X string `json:"url_4x"`
	} `json:"images"`
	Tier       string   `json:"tier"`
	EmoteType  string   `json:"emote_type"`
	EmoteSetId string   `json:"emote_set_id"`
	Format     []string `json:"format"`
	Scale      []string `json:"scale"`
	ThemeMode  []string `json:"theme_mode"`
}

type FollowersResponse struct {
	Total int `json:"total"`
	Data  []struct {
		FromId     string    `json:"from_id"`
		FromLogin  string    `json:"from_login"`
		FromName   string    `json:"from_name"`
		ToId       string    `json:"to_id"`
		ToName     string    `json:"to_name"`
		FollowedAt time.Time `json:"followed_at"`
	} `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type TokenResponse struct {
	AccessToken string   `json:"access_token"`
	ExpiresIn   int      `json:"expires_in"`
	Scope       []string `json:"scope"`
	TokenType   string   `json:"token_type"`
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
	Emotes          EmoteResponse
	IsLive          bool `json:"is_live"`
}

func (user *UserData) GetEmotes(twitchClient TwitchClient) {
	emotes, err := twitchClient.GetEmotes(user.Id)
	if err != nil {
		log.Fatal(err)
	}
	user.Emotes = *emotes
}

type UserDataResponse struct {
	Data       []UserData `json:"data"`
	Pagination struct {
		Cursor string `json:"cursor"`
	} `json:"pagination"`
}

type Stream struct {
	Id        string `json:"id"`
	IsLive    bool   `json:"is_live"`
	ViewCount int    `json:"viewer_count"`
}

type StreamsResponse struct {
	Data []Stream `json:"data"`
}
