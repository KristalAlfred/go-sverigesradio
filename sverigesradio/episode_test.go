package sverigesradio

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetEpisodes(t *testing.T) {
	client := NewClient(http.DefaultClient)

	ID := 3718
	from := time.Date(2021, time.December, 1, 0, 0, 0, 0, &time.Location{})
	to := time.Date(2021, time.December, 1, 6, 0, 0, 0, &time.Location{})

	resp, err := client.Episode.ListEpisodes(context.Background(), &EpisodesOptions{
		ProgramID: &ID,
		FromDate:  &from,
		ToDate:    &to,
		GeneralOptions: GeneralOptions{
			Pagination: false,
			Format:     JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in ListEpisodes(), got error: %v", err)
	}

	assert.Equal(
		t,
		1,
		len(resp.Episodes),
		"There should only be a single episode in the supplied timeslot",
	)

	assert.Equal(
		t,
		"Tankesmedjans Julkalender: Lucka 1",
		*resp.Episodes[0].Title,
		"The title of the episode should be 'Tankesmedjans Julkalender: Lucka 1'",
	)
}

func TestSearchEpisode(t *testing.T) {
	client := NewClient(http.DefaultClient)

	query := "Linnea Henriksson vs. Sandro Cavazza"

	resp, err := client.Episode.SearchEpisode(context.Background(), &EpisodeSearchOptions{
		Query: &query,
		GeneralOptions: GeneralOptions{
			Format:     JSON,
			Pagination: false,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in SearchEpisode(), got error: %v", err)
	}

	assert.Equal(t, 1, len(resp.Episodes), "There is only 1 episode named exactly 'Linnea Henriksson vs. Sandro Cavazza'")
	assert.Equal(t, "Linnea Henriksson vs. Sandro Cavazza", *resp.Episodes[0].Title, "Title should be 'Linnea Henriksson vs. Sandro Cavazza'")
}

func TestGetEpisode(t *testing.T) {
	client := NewClient(http.DefaultClient)

	ID := 602474

	resp, err := client.Episode.GetEpisode(context.Background(), &EpisodeOptions{
		EpisodeID: &ID,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetEpisode(), got error: %v", err)
	}

	assert.Equal(t, "Historien om John Hron", *resp.Episode.Title, "Episode with ID 602474 should have title 'Historien om John Hron'")
}

func TestGetEpisodeList(t *testing.T) {
	client := NewClient(http.DefaultClient)

	episodeIDs := []int{697028, 681604}

	resp, err := client.Episode.GetEpisodesByID(context.Background(), &EpisodeListOptions{
		EpisodeIDs: &episodeIDs,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetEpisodesByID(), got error: %v", err)
	}

	assert.Equal(t, 2, len(resp.Episodes), "We should get the two episodes with IDs corresponding with what we sent in")
	assert.Equal(t, "Avsnitt 16: Hem, trygga hem", *resp.Episodes[0].Title, "The first returned episode should be 'Avsnitt 16: Hem, trygga hem' of program 'Creepypodden i P3'")
}

func TestLatestEpisode(t *testing.T) {
	client := NewClient(http.DefaultClient)

	ID := 3117

	resp, err := client.Episode.GetLatestEpisode(context.Background(), &LatestEpisodeOptions{
		ProgramID: &ID,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetLatestEpisode(), got error: %v", err)
	}

	assert.Equal(t, "Karlavagnen", *resp.Episode.Program.Name, "The episode should belong to the program 'Karlavagnen' (ID 3117)")
}

func TestGetEpisodeByGroup(t *testing.T) {
	client := NewClient(http.DefaultClient)

	ID := 23037

	resp, err := client.Episode.GetEpisodesByGroup(context.Background(), &EpisodeGroupOptions{
		GroupID: &ID,
		GeneralOptions: GeneralOptions{
			Format:     JSON,
			Pagination: false,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetEpisodesByGroup(), got error: %v", err)
	}

	for _, episode := range resp.Episodes {
		assert.Equal(t, *episode.Program.Name, "P1 Dokumentär", "All episodes in category 23027 should be from program P1 Dokumentär")
	}
}
