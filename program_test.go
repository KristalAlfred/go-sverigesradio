package sverigesradio

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestListAllPrograms(t *testing.T) {
	client := NewClient(http.DefaultClient)
	programs, err := client.Program.GetAllPrograms(context.Background(), &ProgramOptions{
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
	programCategories, err := client.Program.ListAllProgramCategories(context.Background(), &GeneralOptions{
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
	broadcasts, err := client.Program.GetAllBroadcasts(context.Background(), &BroadcastOptions{
		ProgramID: 3718,
		GeneralOptions: GeneralOptions{
			Format: JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occurred in GetProgramCategoryByID(), got error: %v", err)
	}
	for _, broadcast := range broadcasts {
		fmt.Println(broadcast.Title)
	}

	t.Errorf("IDSDSD")
}
