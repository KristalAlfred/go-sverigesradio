package sverigesradio

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

func TestListAllPrograms(t *testing.T) {
	client := NewClient(http.DefaultClient)
	programs, err := client.Program.ListAllPrograms(context.Background(), &ProgramOptions{
		GeneralOptions: GeneralOptions{
			Pagination: false,
			Format:     JSON,
		},
	})
	if err != nil {
		t.Errorf("Error occured in ListAllPrograms(), got error: %v", err)
	}
	for _, program := range programs {
		fmt.Println(program.Name)
	}

	t.Error("heyo")
}
