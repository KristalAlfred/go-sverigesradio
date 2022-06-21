package sverigesradio

import (
	"context"
	"path"
	"strconv"
)

const (
	programEndpoint          = "programs"
	programCategoryEndpoint  = "programcategories"
	programBroadcastEndpoint = "broadcasts"
	programPodfileEndpoint   = "podfiles"
)

// This service implements the methods described in this section
// of the official documentation:
// https://api.sr.se/api/documentation/v2/metoder/program.html
type ProgramService service

// Program represents a Sveriges Radio show
type Program struct {
	Description     *string `json:"description,omitempty"`
	Programcategory *struct {
		ID   *int    `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	} `json:"programcategory,omitempty"`
	Email                    *string `json:"email,omitempty"`
	Phone                    *string `json:"phone,omitempty"`
	Programurl               *string `json:"programurl,omitempty"`
	Programslug              *string `json:"programslug,omitempty"`
	Programimage             *string `json:"programimage,omitempty"`
	Programimagetemplate     *string `json:"programimagetemplate,omitempty"`
	Programimagewide         *string `json:"programimagewide,omitempty"`
	Programimagetemplatewide *string `json:"programimagetemplatewide,omitempty"`
	Socialimage              *string `json:"socialimage,omitempty"`
	Socialimagetemplate      *string `json:"socialimagetemplate,omitempty"`
	Socialmediaplatforms     []*struct {
		Platform    *string `json:"platform,omitempty"`
		Platformurl *string `json:"platformurl,omitempty"`
	} `json:"socialmediaplatforms,omitempty"`
	Channel *struct {
		ID   *int    `json:"id,omitempty"`
		Name *string `json:"name,omitempty"`
	} `json:"channel,omitempty"`
	Archived          *bool   `json:"archived,omitempty"`
	Hasondemand       *bool   `json:"hasondemand,omitempty"`
	Haspod            *bool   `json:"haspod,omitempty"`
	Responsibleeditor *string `json:"responsibleeditor,omitempty"`
	ID                *int    `json:"id,omitempty"`
	Name              *string `json:"name,omitempty"`
	Broadcastinfo     *string `json:"broadcastinfo,omitempty"`
	Payoff            *string `json:"payoff,omitempty"`
}

type ProgramOptions struct {
	GeneralOptions
	ChannelID         *int  `url:"channelid,omitempty"`
	ProgramCategoryID *int  `url:"programcategoryid,omitempty"`
	IsArchived        *bool `url:"isarchived,omitempty"`
}

type ProgramResponse struct {
	Program    *Program `json:"program,omitempty"`
	Pagination `json:"pagination,omitempty"`
}

type ProgramsResponse struct {
	Programs   []*Program `json:"programs,omitempty"`
	Pagination `json:"pagination,omitempty"`
}

func (s *ProgramService) GetPrograms(ctx context.Context, opt *ProgramOptions) (*ProgramsResponse, error) {
	r, err := addOptions(programEndpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *ProgramsResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *ProgramService) GetProgramByID(ctx context.Context, id int, generalOptions *GeneralOptions) (*ProgramResponse, error) {
	p := path.Join(programEndpoint, strconv.Itoa(id))
	r, err := addOptions(p, generalOptions)
	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *ProgramResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type ProgramCategory struct {
	Id   *int
	Name *string
}

type ProgramCategoriesResponse struct {
	ProgramCategories []*ProgramCategory `json:"programcategories,omitempty"`
	Pagination        `json:"pagination,omitempty"`
}

func (s *ProgramService) ListProgramCategories(ctx context.Context, opt *GeneralOptions) (*ProgramCategoriesResponse, error) {
	r, err := addOptions(programCategoryEndpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *ProgramCategoriesResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type ProgramCategoryResponse struct {
	ProgramCategory *ProgramCategory `json:"programcategory,omitempty"`
	Pagination      `json:"pagination,omitempty"`
}

func (s *ProgramService) GetProgramCategoryByID(ctx context.Context, id int, opt *GeneralOptions) (*ProgramCategoryResponse, error) {
	p := path.Join(programCategoryEndpoint, strconv.Itoa(id))
	r, err := addOptions(p, opt)
	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *ProgramCategoryResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type Broadcast struct {
	ID               *int    `json:"id,omitempty"`
	Title            *string `json:"title,omitempty"`
	Description      *string `json:"description,omitempty"`
	Broadcastdateutc *string `json:"broadcastdateutc,omitempty"`
	Totalduration    *int    `json:"totalduration,omitempty"`
	Image            *string `json:"image,omitempty"`
	Imagetemplate    *string `json:"imagetemplate,omitempty"`
	Availablestoputc *string `json:"availablestoputc,omitempty"`
	Playlist         *struct {
		Duration       *int    `json:"duration,omitempty"`
		Publishdateutc *string `json:"publishdateutc,omitempty"`
		ID             *int    `json:"id,omitempty"`
		URL            *string `json:"url,omitempty"`
		Statkey        *string `json:"statkey,omitempty"`
	} `json:"playlist,omitempty"`
	Broadcastfiles []*struct {
		Duration       *int    `json:"duration,omitempty"`
		Publishdateutc *string `json:"publishdateutc,omitempty"`
		ID             *int    `json:"id,omitempty"`
		URL            *string `json:"url,omitempty"`
		Statkey        *string `json:"statkey,omitempty"`
	} `json:"broadcastfiles,omitempty"`
}

type BroadcastOptions struct {
	GeneralOptions
	ProgramID *int `url:"programid,omitempty"`
}

type BroadcastsResponse struct {
	Description *string      `json:"description,omitempty"`
	Copyright   *string      `json:"copyright,omitempty"`
	Name        *string      `json:"name,omitempty"`
	Broadcasts  []*Broadcast `json:"broadcasts,omitempty"`
	Pagination
}

func (s *ProgramService) GetProgramBroadcasts(ctx context.Context, opt *BroadcastOptions) (*BroadcastsResponse, error) {
	r, err := addOptions(programBroadcastEndpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *BroadcastsResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type Podfile struct {
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
}

type PodfileOptions struct {
	GeneralOptions
	ProgramID *int `url:"programid,omitempty"`
}

type PodfilesResponse struct {
	Copyright *string    `json:"copyright,omitempty"`
	Podfiles  []*Podfile `json:"podfiles,omitempty"`
	Pagination
}

func (s *ProgramService) GetProgramPodfiles(ctx context.Context, opt *PodfileOptions) (*PodfilesResponse, error) {
	r, err := addOptions(programPodfileEndpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *PodfilesResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

type PodfileResponse struct {
	Copyright *string  `json:"copyright,omitempty"`
	Podfile   *Podfile `json:"podfile,omitempty"`
}

func (s *ProgramService) GetPodfileByID(ctx context.Context, id int, opt *GeneralOptions) (*PodfileResponse, error) {
	endpoint := path.Join(programPodfileEndpoint, strconv.Itoa(id))
	r, err := addOptions(endpoint, opt)
	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *PodfileResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
