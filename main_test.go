package main

import "testing"

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
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			_, err := scrap(tt.site)
			if tt.expected == nil && err != nil {
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
