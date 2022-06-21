package sverigesradio

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetNewsPrograms(t *testing.T) {
	client := NewClient(http.DefaultClient)
	resp, err := client.News.GetNewsPrograms(context.Background(), &GeneralOptions{
		Format: JSON,
	})
	if err != nil {
		t.Errorf("Error occurred in GetProgramCategoryByID(), got error: %v", err)
	}

	assert.Equal(t, "Ekot", *resp.Programs[0].Name, "The first news show should be 'Ekot'")
}

func TestGetNewsEpisodess(t *testing.T) {
	client := NewClient(http.DefaultClient)
	_, err := client.News.GetNewsEpisodes(context.Background(), &GeneralOptions{
		Format: JSON,
	})
	if err != nil {
		t.Errorf("Error occurred in GetProgramCategoryByID(), got error: %v", err)
	}
}
