package sverigesradio

import (
	"context"
	"path"
)

const trafficEndpoint = "traffic"

type TrafficService service

type TrafficArea struct {
	Name                    *string  `json:"name,omitempty"`
	Zoom                    *float32 `json:"zoom,omitempty"`
	Radius                  *float32 `json:"radius,omitempty"`
	Trafficdepartmentunitid *int     `json:"trafficdepartmentunitid,omitempty"`
}

type TrafficAreaOptions struct {
	GeneralOptions
	Latitude  *float64 `url:"latitude,omitempty"`
	Longitude *float64 `url:"longitude,omitempty"`
}

type TrafficAreaResponse struct {
	Copyright *string      `json:"copyright,omitempty"`
	Area      *TrafficArea `json:"area,omitempty"`
}

type TrafficAreasResponse struct {
	Copyright *string        `json:"copyright,omitempty"`
	Areas     []*TrafficArea `json:"areas,omitempty"`
}

func (s *TrafficService) GetArea(ctx context.Context, opt *TrafficAreaOptions) (*TrafficAreaResponse, error) {
	endpoint := path.Join(trafficEndpoint, "areas")
	r, err := addOptions(endpoint, opt)
	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *TrafficAreaResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *TrafficService) GetAreas(ctx context.Context, opt *GeneralOptions) (*TrafficAreasResponse, error) {
	endpoint := path.Join(trafficEndpoint, "areas")
	r, err := addOptions(endpoint, opt)
	req, err := s.client.NewRequest("GET", r, nil)
	if err != nil {
		return nil, err
	}

	var resp *TrafficAreasResponse
	if _, err := s.client.Do(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
