package sverigesradio

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetCurrentlyPlayingSongs(t *testing.T) {
	client := NewClient(http.DefaultClient)
	playlist, err := client.Music.GetCurrentlyPlayingSongs(context.Background(), &ChannelOptions{
		ChannelId: 2576,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetCurrentlyPlayingSongs(), got error: %v", err)
	}
	fmt.Println(playlist)
	t.Errorf("test..")

	//	assert.Equal(t, , "Karlavagnen", "Podfile with ID 4126279 should be from program Karlavagnen")
}
