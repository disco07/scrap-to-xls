package main

import (
	"errors"
	"testing"
)

func TestScrap(t *testing.T) {
	tests := []struct {
		description string
		site        string
		expected    error
	}{
		{
			description: "Download data from page",
			site:        "https://a3f.fr/fr/annuaire_temp.php?all&depart_annuaire",
			expected:    nil,
		},
		{
			description: "Download data from page not working, invalid url",
			site:        "",
			expected:    errors.New("unsupported protocol scheme"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			_, err := scrap(tt.site)
			if tt.expected == nil && err != tt.expected {
				t.Errorf("got %v want %v", err.Error(), tt.expected)
			}

			if tt.expected != nil && err.Error() != tt.expected.Error() {
				t.Errorf("got %v want %v", err.Error(), tt.expected)
			}
		})
	}
}

func TestExcelizeData(t *testing.T) {
	d := []*Data{
		{
			Name:    "Drissa KONE",
			Title:   "Software developer",
			Company: "Freelance",
		},
	}
	tests := []struct {
		description string
		data        []*Data
		expected    error
	}{
		{
			description: "create file xls with data",
			data:        d,
			expected:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			err := excelizeData(tt.data)
			if tt.expected == nil && err != nil {
				t.Errorf("got %v want %v", err.Error(), tt.expected)
			}
		})
	}
}

func TestRunMain(t *testing.T) {
	main()
}
