package sverigesradio

import (
	"net/http"
	"net/url"
)

const (
	BaseURL = "https://api.sr.se/api/v2/"
)

type Client struct {
	BaseURL *url.URL
	client  *http.Client

	common service

	// TODO: Uncomment when these exist
	// Audio          *AudioService
	// Channel       *ChannelService
	// Episode        *EpisodeService
	// Extrabroadcast *ExtrabroadcastService
	// Group          *GroupService
	// Music          *MusicService
	// News           *NewsService
	// Program        *ProgramService
	// Tableau        *TableauService
	// Toplist        *ToplistService
	// Traffic        *TrafficService
}

type service struct {
	client *Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	baseURL, _ := url.Parse(BaseURL)

	c := &Client{
		client:  httpClient,
		BaseURL: baseURL,
	}
	c.common.client = c
	// TODO: Uncomment these when they exist
	// c.Audio = (*&AudioService)(&c.common)
	// c.Channel = (*&ChannelService)(&c.common)
	// c.Episode = (*&EpisodeService)(&c.common)
	// c.Extrabroadcast = (*&ExtrabroadcastService)(&c.common)
	// c.Group = (*&GroupService)(&c.common)
	// c.Music = (*&MusicService)(&c.common)
	// c.News = (*&NewsService)(&c.common)
	// c.Program = (*&ProgramService)(&c.common)
	// c.Tableau = (*&TableauService)(&c.common)
	// c.Toplist = (*&ToplistService)(&c.common)
	// c.Traffic = (*&TrafficService)(&c.common)
	return c
}
