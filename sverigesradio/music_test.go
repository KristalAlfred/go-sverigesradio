package sverigesradio

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestGetCurrentlyPlayingSongs(t *testing.T) {
	client := NewClient(http.DefaultClient)

	ID := 2576

	playlist, err := client.Music.GetCurrentlyPlayingSongs(context.Background(), &ChannelOptions{
		ChannelID: &ID,
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

func TestGetSongsByChannelID(t *testing.T) {
	client := NewClient(http.DefaultClient)

	ID := 3718
	start := time.Date(2017, 1, 1, 0, 0, 0, 0, time.Now().Location())
	end := time.Date(2017, 5, 1, 1, 0, 0, 0, time.Now().Location())

	songs, err := client.Music.GetSongsByChannelID(context.Background(), &SongsOptions{
		ChannelID: &ID,
		StartDate: &start,
		EndDate:   &end,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetCurrentlyPlayingSongs(), got error: %v", err)
	}

	for _, song := range songs {
		fmt.Println(song.Title)
	}

	t.Errorf("test..")

	//	assert.Equal(t, , "Karlavagnen", "Podfile with ID 4126279 should be from program Karlavagnen")
}

func TestGetSongsByProgramID(t *testing.T) {
	client := NewClient(http.DefaultClient)

	ID := 3718
	start := time.Date(2017, 1, 1, 0, 0, 0, 0, time.Now().Location())
	end := time.Date(2017, 5, 1, 1, 0, 0, 0, time.Now().Location())
	songs, err := client.Music.GetSongsByProgramID(context.Background(), &SongsOptions{
		ChannelID: &ID,
		StartDate: &start,
		EndDate:   &end,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetCurrentlyPlayingSongs(), got error: %v", err)
	}

	for _, song := range songs {
		fmt.Println(song.Title)
	}

	t.Errorf("test..")

	//	assert.Equal(t, , "Karlavagnen", "Podfile with ID 4126279 should be from program Karlavagnen")
}

func TestGetSongsByEpisodeID(t *testing.T) {
	client := NewClient(http.DefaultClient)

	ID := 422962
	start := time.Date(2017, 1, 1, 0, 0, 0, 0, time.Now().Location())
	end := time.Date(2017, 5, 1, 1, 0, 0, 0, time.Now().Location())
	songs, err := client.Music.GetSongsByEpisodeID(context.Background(), &SongsOptions{
		ChannelID: &ID,
		StartDate: &start,
		EndDate:   &end,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetCurrentlyPlayingSongs(), got error: %v", err)
	}

	for _, song := range songs {
		fmt.Println(song.Title)
	}

	t.Errorf("test..")

	//	assert.Equal(t, , "Karlavagnen", "Podfile with ID 4126279 should be from program Karlavagnen")
}
