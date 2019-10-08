package models

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

// Goal A longitudinal plan goal
type Goal struct {
	ID             string    `json:"id,omitempty"`
	Consumer       Reference `json:"consumer,omitempty"`
	Progress       string    `json:"progress,omitempty"`
	Description    string    `json:"description"`
	Status         string    `json:"status"`
	CreatedBy      CreatedBy `json:"createdBy"`
	GoalDefinition Reference `json:"goalDefinition,omitempty"`
	CreatedAt      string    `json:"createdAt,omitempty"`
	UpdatedAt      string    `json:"updatedAt,omitempty"`
	Version        int32     `json:"version,omitempty"`
}

// GoalDefinition A longitudinal plan goal definition
type GoalDefinition struct {
	ID          string       `json:"id,omitempty"`
	Text        string       `json:"text"`
	Status      string       `json:"status,omitempty"`
	Categories  []Category   `json:"categories,omitempty"`
	Disciplines []Discipline `json:"disciplines,omitempty"`
	Coding      []Coding     `json:"coding,omitempty"`
}

// GoalDefinitionList A list of goal definitions
type GoalDefinitionList struct {
	Items        []GoalDefinition
	TotalResults int32
	FirstLink    string
	LastLink     string
	PrevLink     string
	NextLink     string
}

// ReadDefinitionsFromCSV Parse Goal Definition objects from the provided CSV file.
// Each line is of the expected format: "Goal Definition Text,Coding System,Coding Code,Coding Display"
func ReadDefinitionsFromCSV(file string) (goalDefinitions []GoalDefinition, err error) {
	csvFile, err := os.Open(file)
	if err != nil {
		err = fmt.Errorf("error opening goal definition csv: %s", err.Error())
		return
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, readErr := reader.Read()
		if readErr == io.EOF {
			break
		} else if readErr != nil {
			err = fmt.Errorf("Problem reading line from csv: %s", err.Error())
			return
		}

		goalDefn := GoalDefinition{
			Text: line[0],
			Coding: []Coding{
				Coding{
					System:  line[1],
					Code:    line[2],
					Display: line[3],
				},
			},
		}

		goalDefinitions = append(goalDefinitions, goalDefn)
	}

	return goalDefinitions, nil
}
