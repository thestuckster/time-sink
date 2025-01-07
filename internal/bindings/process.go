package bindings

import (
	"context"
	"fmt"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/mitchellh/go-ps"
)

type ProcessListDto struct {
	Processes []*ProcessDto `json:"processes"`
}

type ProcessDto struct {
	Name string `json:"name"`
}

type ProcessBinding struct {
	ctx context.Context
}

func NewProcessBinding() *ProcessBinding {
	return &ProcessBinding{}
}

func (pb *ProcessBinding) GetCurrentlyRunningProcesses() *ProcessListDto {
	names := hashset.New()
	procs, err := ps.Processes()
	if err != nil {
		panic(err)
	}

	dtos := make([]*ProcessDto, 0)
	for _, proc := range procs {
		name := proc.Executable()
		if !names.Contains(name) {
			names.Add(name)
		}
	}

	for _, name := range names.Values() {
		dtos = append(dtos, &ProcessDto{
			Name: fmt.Sprintf("%v", name),
		})
	}

	return &ProcessListDto{
		Processes: dtos,
	}
}
