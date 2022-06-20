package sverigesradio

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetExtraBroadcasts(t *testing.T) {
	client := NewClient(http.DefaultClient)
	broadcasts, err := client.Extrabroadcast.GetExtraBroadcasts(context.Background(), &ExtrabroadcastOptions{
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetEpisode(), got error: %v", err)
	}

	for _, broadcast := range broadcasts {
		fmt.Println(broadcast)
	}

	t.Errorf("HEO")
}
