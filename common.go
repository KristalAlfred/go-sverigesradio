package sverigesradio

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

type Pagination struct {
	Page       int `json:"page"`
	Size       int `json:"size"`
	Totalhits  int `json:"totalhits"`
	Totalpages int `json:"totalpages"`
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
