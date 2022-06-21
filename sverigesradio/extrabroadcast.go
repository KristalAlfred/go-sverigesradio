package sverigesradio

import (
	"context"
	"time"
)

const extrabroadcastEndpoint = "extra/broadcasts"

type ExtrabroadcastService service

type Extrabroadcast struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Sport          bool   `json:"sport"`
	Description    string `json:"description"`
	Localstarttime string `json:"localstarttime"`
	Localstoptime  string `json:"localstoptime"`
	Publisher      struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"publisher"`
	Channel struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"channel"`
	Liveaudio struct {
		ID      int    `json:"id"`
		URL     string `json:"url"`
		Statkey string `json:"statkey"`
	} `json:"liveaudio"`
	Mobileliveaudio struct {
		ID      int    `json:"id"`
		URL     string `json:"url"`
		Statkey string `json:"statkey"`
	} `json:"mobileliveaudio"`
}

type ExtrabroadcastOptions struct {
	GeneralOptions
	Date time.Time `url:"date,omitempty"`
}

type ExtrabroadcastsResponse struct {
	Extrabroadcasts []*Extrabroadcast `json:"broadcasts,omitempty"`
	Pagination      `json:"pagination,omitempty"`
}

func (s *ExtrabroadcastService) GetExtrabroadcasts(ctx context.Context, opt *ExtrabroadcastOptions) (*ExtrabroadcastsResponse, error) {
	r, err := addOptions(extrabroadcastEndpoint, opt)
	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *ExtrabroadcastsResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
