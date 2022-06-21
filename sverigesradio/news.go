package sverigesradio

import (
	"context"
	"path"
)

const programNewsEndpoint = "news"

type NewsService service

func (s *NewsService) GetNewsPrograms(ctx context.Context, opt *GeneralOptions) (*ProgramsResponse, error) {
	r, err := addOptions(programNewsEndpoint, opt)
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

func (s *NewsService) GetNewsEpisodes(ctx context.Context, opt *GeneralOptions) (*EpisodesResponse, error) {
	endpoint := path.Join(programNewsEndpoint, "episodes")
	r, err := addOptions(endpoint, opt)
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
