package sverigesradio

import (
	"context"
	"path"
	"time"
)

const musicEndpoint = "playlists"

type MusicService service

type Song struct {
	Title        *string `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	Artist       *string `json:"artist,omitempty"`
	Composer     *string `json:"composer,omitempty"`
	Conductor    *string `json:"conductor,omitempty"`
	Albumname    *string `json:"albumname,omitempty"`
	Recordlabel  *string `json:"recordlabel,omitempty"`
	Lyricist     *string `json:"lyricist,omitempty"`
	Producer     *string `json:"producer,omitempty"`
	Starttimeutc *string `json:"starttimeutc,omitempty"`
	Stoptimeutc  *string `json:"stoptimeutc,omitempty"`
}

type Playlist struct {
	Previoussong *Song    `json:"previoussong,omitempty"`
	Song         *Song    `json:"song,omitempty"`
	NextSong     *Song    `json:"nextsong,omitempty"`
	Channel      *Channel `json:"channel,omitempty"`
}

type ChannelOptions struct {
	GeneralOptions
	ChannelID *int `url:"channelid,omitempty"`
}

type playlistResponse struct {
	Playlist *Playlist `json:"playlist,omitempty"`
}

func (s *MusicService) GetCurrentlyPlayingSongs(ctx context.Context, opt *ChannelOptions) (*Playlist, error) {
	endpoint := path.Join(musicEndpoint, "rightnow")

	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *playlistResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Playlist, nil
}

type SongsOptions struct {
	GeneralOptions
	ChannelID *int       `url:"id,omitempty"`
	StartDate *time.Time `url:"startdatetime,omitempty"`
	EndDate   *time.Time `url:"enddatetime,omitempty"`
}

type SongsResponse struct {
	Songs      []*Song `json:"song,omitempty"`
	Pagination `json:"pagination,omitempty"`
}

func (s *MusicService) GetSongsByChannelID(ctx context.Context, opt *SongsOptions) (*SongsResponse, error) {
	endpoint := path.Join(musicEndpoint, "getplaylistbychannelid")

	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *SongsResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *MusicService) GetSongsByProgramID(ctx context.Context, opt *SongsOptions) (*SongsResponse, error) {
	endpoint := path.Join(musicEndpoint, "getplaylistbyprogramid")

	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *SongsResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *MusicService) GetSongsByEpisodeID(ctx context.Context, opt *SongsOptions) (*SongsResponse, error) {
	endpoint := path.Join(musicEndpoint, "getplaylistbyepisodeid")

	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *SongsResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
