package sverigesradio

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetScheduledEpisodes(t *testing.T) {
	client := NewClient(http.DefaultClient)
	channels, err := client.Tableau.GetLiveSchedule(context.Background(), &ScheduleOptions{
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetProgramCategoryByID(), got error: %v", err)
	}

	for _, channel := range channels {
		if channel.CurrentScheduledEpisode != nil {
			fmt.Println(channel.CurrentScheduledEpisode.Title)
		}
	}

	ID := 163

	channel, err := client.Tableau.GetLiveSchedule(context.Background(), &ScheduleOptions{
		ChannelID: &ID,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetProgramCategoryByID(), got error: %v", err)
	}

	fmt.Println(channel[0].CurrentScheduledEpisode.Title)

	t.Errorf("heyo")

	// assert.Equal(t, podfile.Program.Name, "Karlavagnen", "Podfile with ID 4126279 should be from program Karlavagnen")
}
