package sverigesradio

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetArea(t *testing.T) {
	client := NewClient(http.DefaultClient)

	lat, long := 60.0, 12.0
	resp, err := client.Traffic.GetArea(context.Background(), &TrafficAreaOptions{
		Latitude:  &lat,
		Longitude: &long,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetProgramCategoryByID(), got error: %v", err)
	}

	assert.Equal(t, "Värmland", *resp.Area.Name, "The area for the supplied coordinates should be Värmland")
}

func TestGetAreas(t *testing.T) {
	client := NewClient(http.DefaultClient)
	resp, err := client.Traffic.GetAreas(context.Background(), &GeneralOptions{
		Format: JSON,
	})
	if err != nil {
		t.Errorf("Error occurred in GetProgramCategoryByID(), got error: %v", err)
	}

	assert.Equal(t, 25, len(resp.Areas), "There should be 25 traffic areas")
}
