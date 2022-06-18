package sverigesradio

import (
	"context"
	"net/http"
)

// GeneralParameters represents a set of general parameters that
// can be applied to many of the api methods.
type GeneralOptions struct {
	// Activate pagination when set to true
	// Default: true
	Pagination bool `url:"pagination"`

	// Set the maximum number of elements returned
	// Default: 10
	Size int `url:"size,omitempty"`

	// Specifies the desired page number, given Size
	// Default: 1
	Page int `url:"page,omitempty"`

	// Specifies the desired audio quality
	AudioQuality AudioQuality `url:"quality,omitempty"`

	Format Format `url:"format,omitempty"`
}

type AudioQuality string

const (
	Low      AudioQuality = "lo"
	Standard AudioQuality = "normal"
	High     AudioQuality = "high"
)

type Format string

const (
	JSON Format = "json"
)

func getRequest(s *ProgramService, relativeEndpoint string, ctx context.Context, opt interface{}) (*http.Request, error) {
	r, err := addOptions(relativeEndpoint, opt)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
