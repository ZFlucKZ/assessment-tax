package handlers

import "testing"

func TestCalculatePersonalDeductionTax(t *testing.T) {
	tests := []struct {
		name     string
		income   float64
		personal float64
		want     float64
	}{
		{"Income 2,500,000", 2500000.0, 500000.0, 2000000.0},
		{"Income 2,500,000", 2500000.0, 1000000.0, 1500000.0},
		{"Income 2,500,000", 2500000.0, 1500000.0, 1000000.0},
		{"Income 2,500,000", 2500000.0, 2000000.0, 500000.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			totalIncome := calculatePersonalDeductionTax(tt.income, tt.personal)
			if totalIncome != tt.want {
				t.Errorf("calculatePersonalDeductionTax(%v, %v) = %v; want %v", tt.income, tt.personal, totalIncome, tt.want)
			}
		})
	}
}

func TestCalculateDonationDeductionTax(t *testing.T){
	tests := []struct {
		name     string
		income   float64
		donation float64
		want     float64
	}{
		{"Income 2,500,000", 2500000.0, 10000.0, 2490000.0},
		{"Income 2,500,000", 2500000.0, 30000.0, 2470000.0},
		{"Income 2,500,000", 2500000.0, 50000.0, 2450000.0},
		{"Income 2,500,000", 2500000.0, 100000.0, 2400000.0},
		{"Income 2,500,000", 2500000.0, 150000.0, 2400000.0},
		{"Income 2,500,000", 2500000.0, 200000.0, 2400000.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			totalIncome := calculateDonationDeductionTax(tt.income, tt.donation)
			if totalIncome != tt.want {
				t.Errorf("calculateDonationDeductionTax(%v, %v) = %v; want %v", tt.income, tt.donation, totalIncome, tt.want)
			}
		})
	} 
}