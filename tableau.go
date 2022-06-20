package sverigesradio

import (
	"context"
	"path"
)

const (
	tableauEndpoint = "scheduledepisodes"
)

type TableauService service

type ScheduleEpisode struct {
	Episodeid    int    `json:"episodeid"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Starttimeutc string `json:"starttimeutc"`
	Endtimeutc   string `json:"endtimeutc"`
	Program      struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"program"`
	Channel struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"channel"`
	Imageurl         string `json:"imageurl"`
	Imageurltemplate string `json:"imageurltemplate"`
	Subtitle         string `json:"subtitle,omitempty"`
}

type Schedule []*ScheduleEpisode

type ScheduleOptions struct {
	GeneralOptions
	ChannelID int `url:"channelid,omitempty"`
}

type scheduleResponse struct {
	Copyright  string    `json:"copyright,omitempty"`
	Schedule   *Schedule `json:"schedule,omitempty"`
	Pagination `json:"pagination,omitempty"`
}

func (s *TableauService) GetScheduledEpisodes(ctx context.Context, opt *ScheduleOptions) (*Schedule, error) {
	r, err := addOptions(tableauEndpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *scheduleResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Schedule, nil
}

type liveScheduleResponse struct {
	Copyright  string     `json:"copyright,omitempty"`
	Channels   []*Channel `json:"channels,omitempty"`
	Channel    *Channel   `json:"channel,omitempty"`
	Pagination `json:"pagination,omitempty"`
}

func (s *TableauService) GetLiveSchedule(ctx context.Context, opt *ScheduleOptions) ([]*Channel, error) {
	endpoint := path.Join(tableauEndpoint, "rightnow")
	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *liveScheduleResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	if resp.Channels != nil {
		return resp.Channels, nil
	} else {
		return []*Channel{resp.Channel}, nil
	}
}
