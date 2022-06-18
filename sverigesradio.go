package sverigesradio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
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

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	fullUrl, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, fullUrl.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")

	return req, nil
}
