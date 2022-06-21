package sverigesradio

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetScheduledEpisodes(t *testing.T) {
	client := NewClient(http.DefaultClient)
	resp, err := client.Tableau.GetLiveSchedule(context.Background(), &ScheduleOptions{
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetProgramCategoryByID(), got error: %v", err)
	}

	assert.Equal(t, true, len(resp.Channels) > 30, "Listing all channels live schedules should render a schedule for every Sveriges Radio channel, which is at least 30")

	ID := 164

	resp, err = client.Tableau.GetLiveSchedule(context.Background(), &ScheduleOptions{
		ChannelID: &ID,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetProgramCategoryByID(), got error: %v", err)
	}
	assert.Equal(t, "P3", *resp.Channel.Name, "Channel ID 164 belongs to P3")
}
