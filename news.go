package sverigesradio

import (
	"context"
	"path"
)

const programNewsEndpoint = "news"

type NewsService service

func (s *NewsService) GetNewsPrograms(ctx context.Context, opt *GeneralOptions) ([]*Program, error) {
	r, err := addOptions(programNewsEndpoint, opt)
	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *programsResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp.Programs, nil
}

func (s *NewsService) GetNewsEpisodes(ctx context.Context, opt *GeneralOptions) ([]*Episode, error) {
	endpoint := path.Join(programNewsEndpoint, "episodes")
	r, err := addOptions(endpoint, opt)
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
