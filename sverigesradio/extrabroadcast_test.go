package sverigesradio

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetExtrabroadcasts(t *testing.T) {
	client := NewClient(http.DefaultClient)
	resp, err := client.Extrabroadcast.GetExtrabroadcasts(context.Background(), &ExtrabroadcastOptions{
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetExtrabroadcasts(), got error: %v", err)
	}

	assert.Equal(t, "P4 Plus", resp.Extrabroadcasts[0].Name, "P4 Plus should be the first extrabroadcast")
}
