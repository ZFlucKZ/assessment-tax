package handlers

import (
	"math"

	"github.com/ZFlucKZ/assessment-tax/dto"
)

type ProgressiveTax struct {
	minIncome float64
	maxIncome float64
	rate      float64
}

func newProgressiveTax(minIncome float64, maxIncome float64, rate float64) *ProgressiveTax {
	return &ProgressiveTax{
		minIncome: minIncome,
		maxIncome: maxIncome,
		rate:      rate,
	}
}

func (p *ProgressiveTax) calculateTax(income float64) float64 {
	if income < p.minIncome {
		return 0
	}

	if income > p.maxIncome {
		return (p.maxIncome - p.minIncome) * p.rate
	}

	return (income - p.minIncome) * p.rate
}

func CalculateProgressiveTax(taxDetails *dto.Tax) float64 {
	taxes := []*ProgressiveTax{
		newProgressiveTax(0, 150000, 0),
		newProgressiveTax(150000, 500000, 0.1),
		newProgressiveTax(500000, 1000000, 0.15),
		newProgressiveTax(1000000, 2000000, 0.2),
		newProgressiveTax(2000000, math.Inf(1), 0.35),
	}

	// Find Allowance type Personal
	personalAllowance := dto.AllowanceType{}
	for _, allowance := range taxDetails.Allowances {
		if allowance.AllowanceType == "Personal" {
			personalAllowance = allowance
			break
		}
	}

	taxDetails.TotalIncome = CalculatePersonalDeductionTax(taxDetails.TotalIncome, personalAllowance.Amount)

	var totalTax float64
	for _, tax := range taxes {
		totalTax += tax.calculateTax(taxDetails.TotalIncome)
	}

	return totalTax
}