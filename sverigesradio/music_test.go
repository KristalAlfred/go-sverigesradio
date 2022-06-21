package sverigesradio

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentlyPlayingSongs(t *testing.T) {
	client := NewClient(http.DefaultClient)

	ID := 2576

	resp, err := client.Music.GetCurrentlyPlayingSongs(context.Background(), &ChannelOptions{
		ChannelID: &ID,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetCurrentlyPlayingSongs(), got error: %v", err)
	}

	assert.Equal(t, 2576, *resp.Playlist.Channel.ID, "Channel ID should be the same that we requested")
}

func TestGetSongsByChannelID(t *testing.T) {
	client := NewClient(http.DefaultClient)

	ID := 164
	start := time.Date(2017, 1, 1, 0, 13, 0, 0, time.Now().Location())
	end := time.Date(2017, 1, 1, 0, 15, 0, 0, time.Now().Location())

	resp, err := client.Music.GetSongsByChannelID(context.Background(), &SongsOptions{
		ChannelID: &ID,
		StartDate: &start,
		EndDate:   &end,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetSongsByChannelID(), got error: %v", err)
	}

	assert.Equal(t, 1, len(resp.Songs), "Only 1 played at the specified date in the specified channel (P3)")
	assert.Equal(t, "Simple Minds", *resp.Songs[0].Artist, "The artist should be 'Simple Minds'")
}

func TestGetSongsByProgramID(t *testing.T) {
	client := NewClient(http.DefaultClient)

	ID := 3718
	start := time.Date(2017, 1, 2, 16, 0, 0, 0, time.Now().Location())
	end := time.Date(2017, 1, 2, 17, 0, 0, 0, time.Now().Location())
	resp, err := client.Music.GetSongsByProgramID(context.Background(), &SongsOptions{
		ChannelID: &ID,
		StartDate: &start,
		EndDate:   &end,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetSongsByProgramID(), got error: %v", err)
	}

	assert.Equal(t, 8, len(resp.Songs), "There should be 8 songs in the episode during the specified time")
}

func TestGetSongsByEpisodeID(t *testing.T) {
	client := NewClient(http.DefaultClient)

	ID := 422962
	start := time.Date(2017, 1, 1, 0, 0, 0, 0, time.Now().Location())
	end := time.Date(2017, 1, 1, 0, 0, 0, 0, time.Now().Location())
	resp, err := client.Music.GetSongsByEpisodeID(context.Background(), &SongsOptions{
		ChannelID: &ID,
		StartDate: &start,
		EndDate:   &end,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetSongsByEpisodeID(), got error: %v", err)
	}

	assert.Equal(t, 7, len(resp.Songs), "There should be 7 songs in the episode")
}
