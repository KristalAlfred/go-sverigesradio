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
	Description     string `json:"description,omitempty"`
	Programcategory struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"programcategory,omitempty"`
	Email                    string `json:"email,omitempty"`
	Phone                    string `json:"phone,omitempty"`
	Programurl               string `json:"programurl,omitempty"`
	Programslug              string `json:"programslug,omitempty"`
	Programimage             string `json:"programimage,omitempty"`
	Programimagetemplate     string `json:"programimagetemplate,omitempty"`
	Programimagewide         string `json:"programimagewide,omitempty"`
	Programimagetemplatewide string `json:"programimagetemplatewide,omitempty"`
	Socialimage              string `json:"socialimage,omitempty"`
	Socialimagetemplate      string `json:"socialimagetemplate,omitempty"`
	Socialmediaplatforms     []struct {
		Platform    string `json:"platform,omitempty"`
		Platformurl string `json:"platformurl,omitempty"`
	} `json:"socialmediaplatforms,omitempty"`
	Channel struct {
		ID   int    `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"channel,omitempty"`
	Archived          bool   `json:"archived,omitempty"`
	Hasondemand       bool   `json:"hasondemand,omitempty"`
	Haspod            bool   `json:"haspod,omitempty"`
	Responsibleeditor string `json:"responsibleeditor,omitempty"`
	ID                int    `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
	Broadcastinfo     string `json:"broadcastinfo,omitempty"`
	Payoff            string `json:"payoff,omitempty"`
}

type ProgramOptions struct {
	GeneralOptions
	ChannelID         *int  `url:"channelid,omitempty"`
	ProgramCategoryID *int  `url:"programcategoryid,omitempty"`
	IsArchived        *bool `url:"isarchived,omitempty"`
}

type programsResponse struct {
	Copyright string     `json:"copyright"`
	Programs  []*Program `json:"programs"`
}

func (s *ProgramService) GetAllPrograms(ctx context.Context, opt *ProgramOptions) ([]*Program, error) {
	req, err := getRequest(s, programEndpoint, ctx, opt)
	if err != nil {
		return nil, err
	}

	var resp *programsResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Programs, nil
}

type getProgramsResponse struct {
	Copyright string   `json:"copyright"`
	Program   *Program `json:"program"`
}

func (s *ProgramService) GetProgramByID(ctx context.Context, id int, generalOptions *GeneralOptions) (*Program, error) {
	p := path.Join(programEndpoint, strconv.Itoa(id))
	r, err := addOptions(p, generalOptions)
	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *getProgramsResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Program, nil
}

type ProgramCategory struct {
	Id   int
	Name string
}

type programCategoriesResponse struct {
	Copyright         string             `json:"copyright"`
	ProgramCategories []*ProgramCategory `json:"programcategories"`
}

func (s *ProgramService) ListAllProgramCategories(ctx context.Context, opt *GeneralOptions) ([]*ProgramCategory, error) {
	req, err := getRequest(s, programCategoryEndpoint, ctx, &opt)
	if err != nil {
		return nil, err
	}

	var resp *programCategoriesResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.ProgramCategories, nil
}

type programCategoryResponse struct {
	Copyright       string           `json:"copyright"`
	ProgramCategory *ProgramCategory `json:"programcategory"`
}

func (s *ProgramService) GetProgramCategoryByID(ctx context.Context, id int, opt *GeneralOptions) (*ProgramCategory, error) {
	p := path.Join(programCategoryEndpoint, strconv.Itoa(id))
	r, err := addOptions(p, opt)
	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *programCategoryResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.ProgramCategory, nil
}

type Broadcast struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	Broadcastdateutc string `json:"broadcastdateutc"`
	Totalduration    int    `json:"totalduration"`
	Image            string `json:"image"`
	Imagetemplate    string `json:"imagetemplate"`
	Availablestoputc string `json:"availablestoputc"`
	Playlist         struct {
		Duration       int    `json:"duration"`
		Publishdateutc string `json:"publishdateutc"`
		ID             int    `json:"id"`
		URL            string `json:"url"`
		Statkey        string `json:"statkey"`
	} `json:"playlist"`
	Broadcastfiles []struct {
		Duration       int    `json:"duration"`
		Publishdateutc string `json:"publishdateutc"`
		ID             int    `json:"id"`
		URL            string `json:"url"`
		Statkey        string `json:"statkey"`
	} `json:"broadcastfiles"`
}

type BroadcastOptions struct {
	GeneralOptions
	ProgramID int `url:"programid,omitempty"`
}

type broadcastsResponse struct {
	Description string       `json:"description"`
	Copyright   string       `json:"copyright"`
	Name        string       `json:"name"`
	Broadcasts  []*Broadcast `json:"broadcasts"`
	Pagination
}

func (s *ProgramService) GetAllBroadcasts(ctx context.Context, opt *BroadcastOptions) ([]*Broadcast, error) {
	req, err := getRequest(s, programBroadcastEndpoint, ctx, opt)
	if err != nil {
		return nil, err
	}

	var resp *broadcastsResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Broadcasts, nil
}
