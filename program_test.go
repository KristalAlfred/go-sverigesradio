package sverigesradio

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestListAllPrograms(t *testing.T) {
	client := NewClient(http.DefaultClient)
	programs, err := client.Program.ListAllPrograms(context.Background(), &ListProgramOptions{
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
	program, err := client.Program.FindProgramByID(context.Background(), 1120, &GeneralOptions{
		Format: JSON,
	})
	if err != nil {
		t.Errorf("Error occured in FindProgram(), got error: %v", err)
	}
	fmt.Println(program, err)
	t.Errorf("hsajdaksd")
}
