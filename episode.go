package sverigesradio

import (
	"context"
	"path"
	"time"
)

type EpisodeService service

const (
	episodeEndpoint = "episodes"
)

type Episode struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Program     struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"program"`
	Audiopreference   string `json:"audiopreference"`
	Audiopriority     string `json:"audiopriority"`
	Audiopresentation string `json:"audiopresentation"`
	Publishdateutc    string `json:"publishdateutc"`
	Imageurl          string `json:"imageurl"`
	Imageurltemplate  string `json:"imageurltemplate"`
	Broadcasttime     struct {
		Starttimeutc string `json:"starttimeutc"`
		Endtimeutc   string `json:"endtimeutc"`
	} `json:"broadcasttime"`
	Listenpodfile struct {
		Title           string `json:"title"`
		Description     string `json:"description"`
		Filesizeinbytes int    `json:"filesizeinbytes"`
		Program         struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"program"`
		Availablefromutc string `json:"availablefromutc"`
		Duration         int    `json:"duration"`
		Publishdateutc   string `json:"publishdateutc"`
		ID               int    `json:"id"`
		URL              string `json:"url"`
		Statkey          string `json:"statkey"`
	} `json:"listenpodfile"`
	Downloadpodfile struct {
		Title           string `json:"title"`
		Description     string `json:"description"`
		Filesizeinbytes int    `json:"filesizeinbytes"`
		Program         struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"program"`
		Availablefromutc string `json:"availablefromutc"`
		Duration         int    `json:"duration"`
		Publishdateutc   string `json:"publishdateutc"`
		ID               int    `json:"id"`
		URL              string `json:"url"`
		Statkey          string `json:"statkey"`
	} `json:"downloadpodfile"`
}

type EpisodeOptions struct {
	GeneralOptions
	ProgramID int       `url:"programid"`
	FromDate  time.Time `url:"fromdate,omitempty"`
	ToDate    time.Time `url:"todate,omitempty"`
}

type episodeResponse struct {
	Copyright string `json:"copyright"`
	Episodes  []*Episode
	Pagination
}

func (s *EpisodeService) GetEpisodes(ctx context.Context, opt *EpisodeOptions) ([]*Episode, error) {
	endpoint := path.Join(episodeEndpoint, "index")
	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *episodeResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Episodes, nil
}

type EpisodeSearchOptions struct {
	GeneralOptions
	Query     string
	ChannelID int
	ProgramID int
}

type episodeSearchResponse struct {
	Copyright  string     `json:"copyright,omitempty"`
	Episodes   []*Episode `json:"episodes,omitempty"`
	Pagination `json:"pagination,omitempty"`
}

// Searches for episodes matching a query. The API only allows
// a maximum return size of 25 episodes
func (s *EpisodeService) SearchEpisode(ctx context.Context, opt *EpisodeSearchOptions) ([]*Episode, error) {
	endpoint := path.Join(episodeEndpoint, "search")
	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *episodeResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Episodes, nil
}
