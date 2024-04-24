package handlers

import "testing"

func TestCalculateTaxWithWht(t *testing.T) {
	test_cases := []struct {
		name string
		tax  float64
		wht  float64
		want float64
	}{
		{"Tax 0, WHT 0", 0.0, 0.0, 0.0},
		{"Tax 0, WHT 1000", 0.0, 1000.0, -1000.0},
		{"Tax 1000, WHT 0", 1000.0, 0.0, 1000.0},
		{"Tax 1000, WHT 1000", 1000.0, 1000.0, 0.0},
		{"Tax 1000, WHT 2000", 1000.0, 2000.0, -1000.0},
	}

	for _, tt := range test_cases {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateTaxWithWht(tt.tax, tt.wht)
			if got != tt.want {
				t.Errorf("calculateTaxWithWht(%v, %v) = %v; want %v", tt.tax, tt.wht, got, tt.want)
			}
		})
	}
}