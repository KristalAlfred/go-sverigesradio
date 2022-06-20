package sverigesradio

import (
	"context"
	"time"
)

const extrabroadcastEndpoint = "extra/broadcasts"

type ExtrabroadcastService service

type ExtrabroadcastOptions struct {
	GeneralOptions
	Date time.Time `url:"date,omitempty"`
}

func (s *ExtrabroadcastService) GetExtraBroadcasts(ctx context.Context, opt *ExtrabroadcastOptions) ([]*Broadcast, error) {
	r, err := addOptions(extrabroadcastEndpoint, opt)
	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *broadcastsResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Broadcasts, nil
}
