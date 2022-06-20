package sverigesradio

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetEpisodes(t *testing.T) {
	client := NewClient(http.DefaultClient)

	ID := 3718
	from := time.Date(2021, time.December, 1, 0, 0, 0, 0, &time.Location{})
	to := time.Date(2022, time.January, 1, 0, 0, 0, 0, &time.Location{})
	episodes, err := client.Episode.GetEpisodes(context.Background(), &EpisodesOptions{
		ProgramID: &ID,
		FromDate:  &from,
		ToDate:    &to,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetEpisodes(), got error: %v", err)
	}
	for _, episode := range episodes {
		fmt.Println(episode.Title)
	}

	t.Errorf("IDSDSD")
}

func TestSearchEpisode(t *testing.T) {
	client := NewClient(http.DefaultClient)

	query := "tankesmedjan"

	episodes, err := client.Episode.SearchEpisode(context.Background(), &EpisodeSearchOptions{
		Query: &query,
		GeneralOptions: GeneralOptions{
			Format:     JSON,
			Pagination: false,
			Size:       30,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in SearchEpisode(), got error: %v", err)
	}
	for _, episode := range episodes {
		fmt.Println(episode.Title)
	}

	t.Errorf("IDSDSD")
}

func TestGetEpisode(t *testing.T) {
	client := NewClient(http.DefaultClient)

	ID := 602474

	episode, err := client.Episode.GetEpisode(context.Background(), &EpisodeOptions{
		EpisodeID: &ID,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetEpisode(), got error: %v", err)
	}
	fmt.Println(episode.Title)

	t.Errorf("IDSDSD")
}

func TestGetEpisodeList(t *testing.T) {
	client := NewClient(http.DefaultClient)

	episodeIDs := []int{697028, 681604}

	episodes, err := client.Episode.GetEpisodeList(context.Background(), &EpisodeListOptions{
		EpisodeIDs: &episodeIDs,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetEpisode(), got error: %v", err)
	}
	for _, episode := range episodes {
		fmt.Println(episode.Program.Name)
	}

	t.Errorf("IDSDSD")
}

func TestLatestEpisode(t *testing.T) {
	client := NewClient(http.DefaultClient)

	ID := 3117

	episode, err := client.Episode.GetLatestEpisode(context.Background(), &LatestEpisodeOptions{
		ProgramID: &ID,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetEpisode(), got error: %v", err)
	}
	fmt.Println(episode.Program.Name)

	t.Errorf("IDSDSD")
}

func TestGetEpisodeByGroup(t *testing.T) {
	client := NewClient(http.DefaultClient)

	ID := 23037

	episodes, err := client.Episode.GetEpisodesByGroup(context.Background(), &EpisodeGroupOptions{
		GroupID: &ID,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetEpisode(), got error: %v", err)
	}
	for _, episode := range episodes {
		assert.Equal(t, episode.Program.Name, "P1 Dokumentär", "All episodes in category 23027 should be from program P1 Dokumentär")
	}
}
