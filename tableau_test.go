package sverigesradio

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetScheduledEpisodes(t *testing.T) {
	client := NewClient(http.DefaultClient)
	schedule, err := client.Tableau.GetScheduledEpisodes(context.Background(), &ScheduleOptions{
		ChannelID: 164,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetProgramCategoryByID(), got error: %v", err)
	}

	for _, scheduleEpisode := range *schedule {
		fmt.Println(scheduleEpisode.Title)
	}

	t.Errorf("heyo")

	// assert.Equal(t, podfile.Program.Name, "Karlavagnen", "Podfile with ID 4126279 should be from program Karlavagnen")

}
