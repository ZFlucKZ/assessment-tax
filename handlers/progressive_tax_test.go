package handlers

import (
	"testing"
)

func TestCalculateTax(t *testing.T) {
	tests := []struct {
		name   string
		income float64
		want   float64
	}{
		{"Income -10", -10.0, 0.0},
		{"Income 60,000", 60000.0, 0.0},
		{"Income 150,000", 150000.0, 0.0},
		{"Income 300,000", 300000.0, 15000.0},
		{"Income 500,000", 500000.0, 35000.0},
		{"Income 750,000", 750000.0, 72500.0},
		{"Income 1,000,000", 1000000.0, 110000.0},
		{"Income 1,500,000", 1500000.0, 210000.0},
		{"Income 2,000,000", 2000000.0, 310000.0},
		{"Income 2,500,000", 2500000.0, 485000.0},
	}

	taxes := initProgressiveTax()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var totalTax float64
			
			for _, tax := range taxes {
				tax, _ := tax.calculateTax(tt.income)
				totalTax += tax
			}

			if totalTax != tt.want {
				t.Errorf("calculateTax(%v) = %v; want %v", tt.income, totalTax, tt.want)
			}
		})
	}
}

func TestTaxLevel(t *testing.T){
	taxLevelStr := []string{"0-150,000", "150,001-500,000", "500,001-1,000,000", "1,000,001-2,000,000", "2,000,001 ขึ้นไป"}

	tests := []struct {
		name   string
		income float64
		want   string
	}{
		{"Income -10", -10.0, "0-150,000"},
		{"Income 60,000", 60000.0, "0-150,000"},
		{"Income 150,000", 150000.0, "0-150,000"},
		{"Income 300,000", 300000.0, "150,001-500,000"},
		{"Income 500,000", 500000.0, "500,001-1,000,000"},
		{"Income 750,000", 750000.0, "500,001-1,000,000"},
		{"Income 1,000,000", 1000000.0, "1,000,001-2,000,000"},
		{"Income 1,500,000", 1500000.0, "1,000,001-2,000,000"},
		{"Income 2,000,000", 2000000.0, "2,000,001 ขึ้นไป"},
		{"Income 2,500,000", 2500000.0, "2,000,001 ขึ้นไป"},
	}

	taxes := initProgressiveTax()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, tax := range taxes {
				_, taxLevel := tax.calculateTax(tt.income)
				if taxLevel.Level != taxLevelStr[i] {
					t.Errorf("calculateTax(%v) = %v; want %v", tt.income, taxLevel.Level, tt.want)
				}
			}
		})
	}
}