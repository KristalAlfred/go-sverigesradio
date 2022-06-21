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

// Represents a single episode of a program
type Episode struct {
	ID          *int    `json:"id,omitempty"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	URL         *string `json:"url,omitempty"`
	Program     *struct {
		ID   *int    `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	} `json:"program,omitempty"`
	Audiopreference   *string `json:"audiopreference,omitempty"`
	Audiopriority     *string `json:"audiopriority,omitempty"`
	Audiopresentation *string `json:"audiopresentation,omitempty"`
	Publishdateutc    *string `json:"publishdateutc,omitempty"`
	Imageurl          *string `json:"imageurl,omitempty"`
	Imageurltemplate  *string `json:"imageurltemplate,omitempty"`
	Broadcasttime     *struct {
		Starttimeutc *string `json:"starttimeutc,omitempty"`
		Endtimeutc   *string `json:"endtimeutc,omitempty"`
	} `json:"broadcasttime,omitempty"`
	Listenpodfile *struct {
		Title           *string `json:"title,omitempty"`
		Description     *string `json:"description,omitempty"`
		Filesizeinbytes *int    `json:"filesizeinbytes,omitempty"`
		Program         *struct {
			ID   *int    `json:"id,omitempty"`
			Name *string `json:"name,omitempty"`
		} `json:"program,omitempty"`
		Availablefromutc *string `json:"availablefromutc,omitempty"`
		Duration         *int    `json:"duration,omitempty"`
		Publishdateutc   *string `json:"publishdateutc,omitempty"`
		ID               *int    `json:"id,omitempty"`
		URL              *string `json:"url,omitempty"`
		Statkey          *string `json:"statkey,omitempty"`
	} `json:"listenpodfile,omitempty"`
	Downloadpodfile *struct {
		Title           *string `json:"title,omitempty"`
		Description     *string `json:"description,omitempty"`
		Filesizeinbytes *int    `json:"filesizeinbytes,omitempty"`
		Program         *struct {
			ID   *int    `json:"id,omitempty"`
			Name *string `json:"name,omitempty"`
		} `json:"program,omitempty"`
		Availablefromutc *string `json:"availablefromutc,omitempty"`
		Duration         *int    `json:"duration,omitempty"`
		Publishdateutc   *string `json:"publishdateutc,omitempty"`
		ID               *int    `json:"id,omitempty"`
		URL              *string `json:"url,omitempty"`
		Statkey          *string `json:"statkey,omitempty"`
	} `json:"downloadpodfile,omitempty"`
}

type EpisodeOptions struct {
	GeneralOptions
	EpisodeID *int `url:"id,omitempty"`
}

type EpisodeResponse struct {
	Episode    *Episode `json:"episode,omitempty"`
	Pagination `json:"pagination,omitempty"`
}

type EpisodesOptions struct {
	GeneralOptions
	ProgramID *int       `url:"programid,omitempty"`
	FromDate  *time.Time `url:"fromdate,omitempty"`
	ToDate    *time.Time `url:"todate,omitempty"`
}

type EpisodesResponse struct {
	Episodes   []*Episode `json:"episodes,omitempty"`
	Pagination `json:"pagination,omitempty"`
}

func (s *EpisodeService) ListEpisodes(ctx context.Context, opt *EpisodesOptions) (*EpisodesResponse, error) {
	endpoint := path.Join(episodeEndpoint, "index")
	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *EpisodesResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type EpisodeSearchOptions struct {
	GeneralOptions
	Query     *string `url:"query,omitempty"`
	ChannelID *int    `url:"channelid,omitempty"`
	ProgramID *int    `url:"programid,omitempty"`
}

// Searches for episodes matching a query. The API only allows
// a maximum return size of 25 episodes
func (s *EpisodeService) SearchEpisode(ctx context.Context, opt *EpisodeSearchOptions) (*EpisodesResponse, error) {
	endpoint := path.Join(episodeEndpoint, "search")
	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *EpisodesResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *EpisodeService) GetEpisode(ctx context.Context, opt *EpisodeOptions) (*EpisodeResponse, error) {
	endpoint := path.Join(episodeEndpoint, "get")

	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *EpisodeResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type EpisodeListOptions struct {
	GeneralOptions
	EpisodeIDs *[]int `url:"ids,comma,omitempty"`
}

type EpisodeListResponse struct {
	Episodes   []*Episode `json:"episodes,omitempty"`
	Pagination `json:"pagination,omitempty"`
}

func (s *EpisodeService) GetEpisodesByID(ctx context.Context, opt *EpisodeListOptions) (*EpisodeListResponse, error) {
	endpoint := path.Join(episodeEndpoint, "getlist")

	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *EpisodeListResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type LatestEpisodeOptions struct {
	GeneralOptions
	ProgramID *int `url:"programid,omitempty"`
}

type LatestEpisodeResponse struct {
	Episode    *Episode `json:"episode,omitempty"`
	Pagination `json:"pagination,omitempty"`
}

func (s *EpisodeService) GetLatestEpisode(ctx context.Context, opt *LatestEpisodeOptions) (*LatestEpisodeResponse, error) {
	endpoint := path.Join(episodeEndpoint, "getlatest")

	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *LatestEpisodeResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type EpisodeGroupOptions struct {
	GeneralOptions
	GroupID *int `url:"id,omitempty"`
}

type EpisodeGroup struct {
	ID          *int       `json:"id,omitempty"`
	Title       *string    `json:"title,omitempty"`
	Description *string    `json:"description,omitempty"`
	Episodes    []*Episode `json:"episodes,omitempty"`
}

type EpisodeGroupResponse struct {
	EpisodeGroup `json:"episodegroup,omitempty"`
	Pagination   `json:"pagination,omitempty"`
}

func (s *EpisodeService) GetEpisodesByGroup(ctx context.Context, opt *EpisodeGroupOptions) (*EpisodeGroupResponse, error) {
	endpoint := path.Join(episodeEndpoint, "group")

	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *EpisodeGroupResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
