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
