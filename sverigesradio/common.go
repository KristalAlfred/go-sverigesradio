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

type PaginationOptions struct {
	// Activate pagination when set to true
	// Default: true
	Page int `json:"page,omitempty"`

	// Set the maximum number of elements returned
	// Default: 10
	Size         int     `json:"size,omitempty"`
	Totalhits    int     `json:"totalhits,omitempty"`
	Totalpages   int     `json:"totalpages,omitempty"`
	NextPage     *string `json:"nextpage,omitempty"`
	PreviousPage *string `json:"previouspage,omitempty"`
}

type PaginationResult struct {
	Page         *int    `json:"page,omitempty"`
	Size         *int    `json:"size,omitempty"`
	TotalHits    *int    `json:"totalhits,omitempty"`
	TotalPages   *int    `json:"totalpages,omitempty"`
	NextPage     *string `json:"nextpage,omitempty"`
	PreviousPage *string `json:"previouspage,omitempty"`
}

type Channel struct {
	Id                       *int     `json:"id,omitempty"`
	Name                     *string  `json:"name,omitempty"`
	PreviousScheduledEpisode *Episode `json:"previousscheduledepisode,omitempty"`
	CurrentScheduledEpisode  *Episode `json:"currentscheduledepisode,omitempty"`
	NextScheduledEpisode     *Episode `json:"nextscheduledepisode,omitempty"`
}
