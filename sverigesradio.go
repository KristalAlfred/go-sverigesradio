package sverigesradio

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL = "https://api.sr.se/api/v2/"
)

type Client struct {
	BaseURL *url.URL
	client  *http.Client

	common service

	// TODO: Uncomment when these exist
	// Audio          *AudioService
	// Channel       *ChannelService
	Episode *EpisodeService
	// Extrabroadcast *ExtrabroadcastService
	// Group          *GroupService
	// Music          *MusicService
	// News 		  *NewsService
	Program *ProgramService
	// Tableau        *TableauService
	// Toplist        *ToplistService
	// Traffic        *TrafficService
}

type service struct {
	client *Client
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client:  httpClient,
		BaseURL: baseURL,
	}
	c.common.client = c
	// TODO: Uncomment these when they exist
	// c.Audio = (*&AudioService)(&c.common)
	// c.Channel = (*&ChannelService)(&c.common)
	c.Episode = (*EpisodeService)(&c.common)
	// c.Extrabroadcast = (*&ExtrabroadcastService)(&c.common)
	// c.Group = (*&GroupService)(&c.common)
	// c.Music = (*&MusicService)(&c.common)
	// c.News = (*&NewsService)(&c.common)
	c.Program = (*ProgramService)(&c.common)
	// c.Tableau = (*&TableauService)(&c.common)
	// c.Toplist = (*&ToplistService)(&c.common)
	// c.Traffic = (*&TrafficService)(&c.common)
	return c
}

// addOptions adds the parameters in opt as URL query parameters to s. opt
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opt interface{}) (string, error) {
	v := reflect.ValueOf(opt)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opt)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	fullUrl, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, fullUrl.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("Accept", "application/json")

	return req, nil
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		// If we got an error, and the context has been canceled,
		// the context's error is probably more useful.
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}

	defer func() {
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(v)
			if err == io.EOF {
				err = nil
			}
		}
	}

	return resp, err
}
