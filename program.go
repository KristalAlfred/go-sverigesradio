package sverigesradio

import (
	"context"
	"fmt"
)

const (
	programEndpoint = "programs"
)

type ProgramService service

// Program represents a Sveriges Radio show. Note: different from an episode
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
	ChannelID         *int  `json:"channelid,omitempty"`
	ProgramCategoryID *int  `json:"programcategoryid,omitempty"`
	IsArchived        *bool `json:"isarchived,omitempty"`
}

type ListProgramsResponse struct {
	Copyright string     `json:"copyright"`
	Programs  []*Program `json:"programs"`
}

func (s *ProgramService) ListAllPrograms(ctx context.Context, opt *ProgramOptions) ([]*Program, error) {
	r, err := addOptions(programEndpoint, opt)
	if err != nil {
		return nil, err
	}
	r = r + "&format=json&pagination=false"
	fmt.Printf("Printing options: %v\n", opt)

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *ListProgramsResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Programs, nil
}
