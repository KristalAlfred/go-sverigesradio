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

type ListProgramOptions struct {
	GeneralOptions
	ChannelID         *int  `url:"channelid,omitempty"`
	ProgramCategoryID *int  `url:"programcategoryid,omitempty"`
	IsArchived        *bool `url:"isarchived,omitempty"`
}

type ListProgramsResponse struct {
	Copyright string     `json:"copyright"`
	Programs  []*Program `json:"programs"`
}

func (s *ProgramService) ListAllPrograms(ctx context.Context, opt *ListProgramOptions) ([]*Program, error) {
	req, err := getRequest(s, programEndpoint, ctx, opt)
	if err != nil {
		return nil, err
	}

	var resp *ListProgramsResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Programs, nil
}

type FindProgramsResponse struct {
	Copyright string   `json:"copyright"`
	Program   *Program `json:"program"`
}

func (s *ProgramService) FindProgramByID(ctx context.Context, programID int, generalOptions *GeneralOptions) (*Program, error) {
	p := path.Join(programEndpoint, strconv.Itoa(programID))
	r, err := addOptions(p, generalOptions)
	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *FindProgramsResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Program, nil
}

type ProgramCategory struct {
	Id   int
	Name string
}

type ProgramCategoriesResponse struct {
	Copyright         string             `json:"copyright"`
	ProgramCategories []*ProgramCategory `json:"programcategories"`
}

func (s *ProgramService) ListAllProgramCategories(ctx context.Context, opt *GeneralOptions) ([]*ProgramCategory, error) {
	req, err := getRequest(s, programCategoryEndpoint, ctx, &opt)
	if err != nil {
		return nil, err
	}

	var resp *ProgramCategoriesResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.ProgramCategories, nil
}
