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

type EpisodesOptions struct {
	GeneralOptions
	ProgramID *int       `url:"programid,omitempty"`
	FromDate  *time.Time `url:"fromdate,omitempty"`
	ToDate    *time.Time `url:"todate,omitempty"`
}

type episodesResponse struct {
	Copyright *string `json:"copyright,omitempty"`
	Episodes  []*Episode
	Pagination
}

func (s *EpisodeService) GetEpisodes(ctx context.Context, opt *EpisodesOptions) ([]*Episode, error) {
	endpoint := path.Join(episodeEndpoint, "index")
	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *episodesResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Episodes, nil
}

type EpisodeSearchOptions struct {
	GeneralOptions
	Query     *string
	ChannelID *int
	ProgramID *int
}

type episodeSearchResponse struct {
	Copyright  *string    `json:"copyright,omitempty"`
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

	var resp *episodesResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Episodes, nil
}

type EpisodeOptions struct {
	GeneralOptions
	EpisodeID *int `url:"id,omitempty"`
}

type episodeResponse struct {
	Copyright *string  `json:"copyright,omitempty"`
	Episode   *Episode `json:"episode,omitempty"`
}

func (s *EpisodeService) GetEpisode(ctx context.Context, opt *EpisodeOptions) (*Episode, error) {
	endpoint := path.Join(episodeEndpoint, "get")

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
	return resp.Episode, nil
}

type EpisodeListOptions struct {
	GeneralOptions
	EpisodeIDs *[]int `url:"ids,comma,omitempty"`
}

type episodeListResponse struct {
	Copyright *string    `json:"copyright,omitempty"`
	Episodes  []*Episode `json:"episodes,omitempty"`
	Pagination
}

func (s *EpisodeService) GetEpisodeList(ctx context.Context, opt *EpisodeListOptions) ([]*Episode, error) {
	endpoint := path.Join(episodeEndpoint, "getlist")

	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *episodeListResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Episodes, nil
}

type LatestEpisodeOptions struct {
	GeneralOptions
	ProgramID *int `url:"programid,omitempty"`
}

type latestEpisodeResponse struct {
	Copyright *string  `json:"copyright,omitempty"`
	Episode   *Episode `json:"episode,omitempty"`
}

func (s *EpisodeService) GetLatestEpisode(ctx context.Context, opt *LatestEpisodeOptions) (*Episode, error) {
	endpoint := path.Join(episodeEndpoint, "getlatest")

	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *latestEpisodeResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Episode, nil
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

type episodeGroupResponse struct {
	Copyright    *string `json:"copyright,omitempty"`
	EpisodeGroup `json:"episodegroup,omitempty"`
	Pagination
}

func (s *EpisodeService) GetEpisodesByGroup(ctx context.Context, opt *EpisodeGroupOptions) ([]*Episode, error) {
	endpoint := path.Join(episodeEndpoint, "group")

	r, err := addOptions(endpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *episodeGroupResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Episodes, nil
}
