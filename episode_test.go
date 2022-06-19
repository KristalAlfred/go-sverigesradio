package sverigesradio

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestGetEpisodes(t *testing.T) {
	client := NewClient(http.DefaultClient)
	episodes, err := client.Episode.GetEpisodes(context.Background(), &EpisodeOptions{
		ProgramID: 3718,
		FromDate:  time.Date(2021, time.December, 1, 0, 0, 0, 0, &time.Location{}),
		ToDate:    time.Date(2022, time.January, 1, 0, 0, 0, 0, &time.Location{}),
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

func TestSearchEpisode(t *testing.T) {
	client := NewClient(http.DefaultClient)
	episodes, err := client.Episode.SearchEpisode(context.Background(), &EpisodeSearchOptions{
		Query: "tankesmedjan",
		GeneralOptions: GeneralOptions{
			Format:     JSON,
			Pagination: false,
			Size:       30,
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
