package sverigesradio

// Set up some initial constants
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

// GeneralParameters represents a set of general parameters that
// can be applied to almost all of the API methods
type GeneralOptions struct {
	// Activate pagination when set to true
	// Default: true
	Pagination bool `url:"pagination"`

	// Specifies the desired audio quality
	AudioQuality AudioQuality `url:"quality,omitempty"`

	// Specifies the desired API response format. This client library only supports JSON,
	// but the Sveriges Radio API also supports XML & JSONP
	Format Format `url:"format,omitempty"`
}

type Pagination struct {
	// Specifies the desired page number, given Size
	// Default: 1
	Page *int `json:"page,omitempty"`

	// Set the maximum number of elements returned
	// Default: 10
	Size         *int    `json:"size,omitempty"`
	TotalHits    *int    `json:"totalhits,omitempty"`
	TotalPages   *int    `json:"totalpages,omitempty"`
	NextPage     *string `json:"nextpage,omitempty"`
	PreviousPage *string `json:"previouspage,omitempty"`
}

// Represents the
type Channel struct {
	ID                       *int     `json:"id,omitempty"`
	Name                     *string  `json:"name,omitempty"`
	PreviousScheduledEpisode *Episode `json:"previousscheduledepisode,omitempty"`
	CurrentScheduledEpisode  *Episode `json:"currentscheduledepisode,omitempty"`
	NextScheduledEpisode     *Episode `json:"nextscheduledepisode,omitempty"`
}
