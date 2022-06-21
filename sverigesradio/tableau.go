package sverigesradio

import (
	"context"
	"fmt"
	"path"
)

const tableauEndpoint = "scheduledepisodes"

type TableauService service

type ScheduleEpisode struct {
	Episodeid    *int    `json:"episodeid,omitempty"`
	Title        *string `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	Starttimeutc *string `json:"starttimeutc,omitempty"`
	Endtimeutc   *string `json:"endtimeutc,omitempty"`
	Program      *struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"program,omitempty"`
	Channel          *ChannelSchedule `json:"channel,omitempty"`
	Imageurl         *string          `json:"imageurl,omitempty"`
	Imageurltemplate *string          `json:"imageurltemplate,omitempty"`
	Subtitle         *string          `json:"subtitle,omitempty"`
}

type Schedule []*ScheduleEpisode

type ScheduleOptions struct {
	GeneralOptions
	ChannelID *int `url:"channelid,omitempty"`
}

type ScheduleResponse struct {
	Schedule   *Schedule `json:"schedule,omitempty"`
	Pagination `json:"pagination,omitempty"`
}

func (s *TableauService) GetScheduledEpisodes(ctx context.Context, opt *ScheduleOptions) (*ScheduleResponse, error) {
	r, err := addOptions(tableauEndpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(req)

	var resp *ScheduleResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// This represents the API response when querying for
type LiveScheduleResponse struct {
	Channels   []*ChannelSchedule `json:"channels,omitempty"`
	Channel    *ChannelSchedule   `json:"channel,omitempty"`
	Pagination `json:"pagination,omitempty"`
}

func (s *TableauService) GetLiveSchedule(ctx context.Context, opt *ScheduleOptions) (*LiveScheduleResponse, error) {
	endpoint := path.Join(tableauEndpoint, "rightnow")
	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *LiveScheduleResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}
