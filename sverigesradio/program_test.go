package sverigesradio

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPrograms(t *testing.T) {
	client := NewClient(http.DefaultClient)
	resp, err := client.Program.GetPrograms(context.Background(), &ProgramOptions{
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occured in ListAllPrograms(), got error: %v", err)
	}

	assert.Equal(t, true, len(resp.Programs) > 200, "The method should return a large list of programs")
}

func TestGetProgramByID(t *testing.T) {
	client := NewClient(http.DefaultClient)
	id := 1120
	resp, err := client.Program.GetProgramByID(context.Background(), id, &GeneralOptions{
		Format: JSON,
	})
	if err != nil {
		t.Errorf("Error occured in GetProgramByID(), got error: %v", err)
	}

	assert.Equal(t, id, *resp.Program.ID, "ID received should match the input")
}

func TestListProgramCategories(t *testing.T) {
	client := NewClient(http.DefaultClient)
	resp, err := client.Program.ListProgramCategories(context.Background(), &GeneralOptions{
		Format: JSON,
	})
	if err != nil {
		t.Errorf("Error occurred in ListAllProgramCategories(), got error: %v", err)
	}

	assert.Equal(t, "Ekonomi", *resp.ProgramCategories[4].Name, "The fifth category name should be 'Ekonomi'")
}

func TestGetProgramCategoryByID(t *testing.T) {
	client := NewClient(http.DefaultClient)
	resp, err := client.Program.GetProgramCategoryByID(context.Background(), 5, &GeneralOptions{
		Format: JSON,
	})
	if err != nil {
		t.Errorf("Error occurred in GetProgramCategoryByID(), got error: %v", err)
	}

	assert.Equal(t, "Musik", *resp.ProgramCategory.Name, "Category with ID 5 should be 'Musik'")
}

func TestGetBroadcasts(t *testing.T) {
	client := NewClient(http.DefaultClient)

	ID := 3718

	_, err := client.Program.GetProgramBroadcasts(context.Background(), &BroadcastOptions{
		ProgramID: &ID,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetGetAllBroadcasts(), got error: %v", err)
	}
}

func TestGetProgramPodfiles(t *testing.T) {
	client := NewClient(http.DefaultClient)

	ID := 3117

	resp, err := client.Program.GetProgramPodfiles(context.Background(), &PodfileOptions{
		ProgramID: &ID,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetAllProgramPodfiles(), got error: %v", err)
	}

	assert.Equal(t, true, len(resp.Podfiles) > 200, "There should be at least 200 podfiles for Karlavagnen")
}

func TestGetPodfileByID(t *testing.T) {
	client := NewClient(http.DefaultClient)
	resp, err := client.Program.GetPodfileByID(context.Background(), 4126279, &GeneralOptions{
		Format: JSON,
	})
	if err != nil {
		t.Errorf("Error occurred in GetPodfileByID(), got error: %v", err)
	}

	assert.Equal(t, *resp.Podfile.Program.Name, "Karlavagnen", "Podfile with ID 4126279 is a podfile from the show Karlavagnen")
}
