package sverigesradio

import (
	"context"
	"fmt"
	"path"
	"time"
)

const (
	musicEndpoint = "playlists"
)

type MusicService service

type Song struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	Artist       string `json:"artist"`
	Composer     string `json:"composer"`
	Conductor    string `json:"conductor"`
	Albumname    string `json:"albumname"`
	Recordlabel  string `json:"recordlabel"`
	Lyricist     string `json:"lyricist"`
	Producer     string `json:"producer"`
	Starttimeutc string `json:"starttimeutc"`
	Stoptimeutc  string `json:"stoptimeutc"`
}

type Playlist struct {
	Previoussong Song    `json:"previoussong,omitempty"`
	Song         Song    `json:"song,omitempty"`
	NextSong     Song    `json:"nextsong,omitempty"`
	Channel      Channel `json:"channel,omitempty"`
}

type ChannelOptions struct {
	GeneralOptions
	ChannelID int `url:"channelid"`
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
	ID        int       `url:"id,omitempty"`
	StartDate time.Time `url:"startdatetime,omitempty"`
	EndDate   time.Time `url:"enddatetime,omitempty"`
}

type songsResponse struct {
	Songs []*Song `json:"song,omitempty"`
}

func (s *MusicService) GetSongsByChannelID(ctx context.Context, opt *SongsOptions) ([]*Song, error) {
	endpoint := path.Join(musicEndpoint, "getplaylistbychannelid")

	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}
	fmt.Println(r)

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *songsResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Songs, nil
}

func (s *MusicService) GetSongsByProgramID(ctx context.Context, opt *SongsOptions) ([]*Song, error) {
	endpoint := path.Join(musicEndpoint, "getplaylistbyprogramid")

	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}
	fmt.Println(r)

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *songsResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Songs, nil
}

func (s *MusicService) GetSongsByEpisodeID(ctx context.Context, opt *SongsOptions) ([]*Song, error) {
	endpoint := path.Join(musicEndpoint, "getplaylistbyepisodeid")

	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}
	fmt.Println(r)

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *songsResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Songs, nil
}
