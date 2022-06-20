package sverigesradio

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListAllPrograms(t *testing.T) {
	client := NewClient(http.DefaultClient)
	programs, err := client.Program.GetPrograms(context.Background(), &ProgramOptions{
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occured in ListAllPrograms(), got error: %v", err)
	}

	for _, program := range programs {
		fmt.Println(program.Name)
	}

	t.Errorf("heyo")
}

func TestFindProgramByID(t *testing.T) {
	client := NewClient(http.DefaultClient)
	program, err := client.Program.GetProgramByID(context.Background(), 1120, &GeneralOptions{
		Format: JSON,
	})
	if err != nil {
		t.Errorf("Error occured in GetProgramByID(), got error: %v", err)
	}
	fmt.Println(program, err)
	t.Errorf("hsajdaksd")
}

func TestListAllProgramCategories(t *testing.T) {
	client := NewClient(http.DefaultClient)
	programCategories, err := client.Program.ListProgramCategories(context.Background(), &GeneralOptions{
		Format: JSON,
	})
	if err != nil {
		t.Errorf("Error occurred in ListAllProgramCategories(), got error: %v", err)
	}

	for _, category := range programCategories {
		fmt.Println(category.Name)
	}
	t.Errorf("HEYO! :D")
}

func TestGetProgramCategoryByID(t *testing.T) {
	client := NewClient(http.DefaultClient)
	programCategory, err := client.Program.GetProgramCategoryByID(context.Background(), 2, &GeneralOptions{
		Format: JSON,
	})
	if err != nil {
		t.Errorf("Error occurred in GetProgramCategoryByID(), got error: %v", err)
	}
	fmt.Println(programCategory)

	t.Errorf("IDSDSD")
}

func TestGetAllBroadcasts(t *testing.T) {
	client := NewClient(http.DefaultClient)

	ID := 3718

	broadcasts, err := client.Program.GetProgramBroadcasts(context.Background(), &BroadcastOptions{
		ProgramID: &ID,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetProgramCategoryByID(), got error: %v", err)
	}
	for _, broadcast := range broadcasts {
		for _, broadcastfile := range broadcast.Broadcastfiles {
			fmt.Println(broadcastfile.URL)
		}
	}

	t.Errorf("IDSDSD")
}

func TestGetAllProgramPodfiles(t *testing.T) {
	client := NewClient(http.DefaultClient)

	ID := 3117

	podfiles, err := client.Program.GetProgramPodfiles(context.Background(), &PodfileOptions{
		ProgramID: &ID,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetProgramCategoryByID(), got error: %v", err)
	}
	for _, podfile := range podfiles {
		fmt.Println(podfile.Title)
	}

	t.Errorf("IDSDSD")
}

func TestGetPodfileByID(t *testing.T) {
	client := NewClient(http.DefaultClient)
	podfile, err := client.Program.GetPodfileByID(context.Background(), 4126279, &GeneralOptions{
		Format: JSON,
	})
	if err != nil {
		t.Errorf("Error occurred in GetProgramCategoryByID(), got error: %v", err)
	}

	assert.Equal(t, podfile.Program.Name, "Karlavagnen", "Podfile with ID 4126279 should be from program Karlavagnen")
}
