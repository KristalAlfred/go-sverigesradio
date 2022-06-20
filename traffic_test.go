package sverigesradio

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestGetArea(t *testing.T) {
	client := NewClient(http.DefaultClient)
	area, err := client.Traffic.GetArea(context.Background(), &TrafficAreaOptions{
		Latitude:  60.0,
		Longitude: 12.0,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetProgramCategoryByID(), got error: %v", err)
	}

	fmt.Println(area.Name)
	t.Errorf("Yo")
}

func TestGetAreas(t *testing.T) {
	client := NewClient(http.DefaultClient)
	areas, err := client.Traffic.GetAreas(context.Background(), &GeneralOptions{
		Format: JSON,
	})
	if err != nil {
		t.Errorf("Error occurred in GetProgramCategoryByID(), got error: %v", err)
	}

	for _, area := range areas {
		fmt.Println(area.Name)
	}
	t.Errorf("Yo")
}
