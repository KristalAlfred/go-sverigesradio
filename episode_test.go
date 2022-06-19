package sverigesradio

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetEpisodes(t *testing.T) {
	client := NewClient(http.DefaultClient)
	episodes, err := client.Episode.GetEpisodes(context.Background(), &EpisodeOptions{
		ProgramID: 3718,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetProgramCategoryByID(), got error: %v", err)
	}
	for _, episode := range episodes {
		fmt.Println(episode.Title)
	}

	t.Errorf("IDSDSD")
}
