package sverigesradio

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetNewsPrograms(t *testing.T) {
	client := NewClient(http.DefaultClient)
	programs, err := client.News.GetNewsPrograms(context.Background(), &GeneralOptions{
		Format: JSON,
	})
	if err != nil {
		t.Errorf("Error occurred in GetProgramCategoryByID(), got error: %v", err)
	}

	for _, program := range programs {
		fmt.Println(program.ID)
	}

	t.Errorf("Heyoo")

	//assert.Equal(t, podfile.Program.Name, "Karlavagnen", "Podfile with ID 4126279 should be from program Karlavagnen")
}

func TestGetNewsEpisodess(t *testing.T) {
	client := NewClient(http.DefaultClient)
	episodes, err := client.News.GetNewsEpisodes(context.Background(), &GeneralOptions{
		Format: JSON,
	})
	if err != nil {
		t.Errorf("Error occurred in GetProgramCategoryByID(), got error: %v", err)
	}

	for _, episode := range episodes {
		fmt.Println(episode.ID)
	}

	t.Errorf("Heyoo")

	//assert.Equal(t, podfile.Program.Name, "Karlavagnen", "Podfile with ID 4126279 should be from program Karlavagnen")
}
