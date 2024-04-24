package handlers

import (
	"math"
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

func initProgressiveTax() []*ProgressiveTax {
	return []*ProgressiveTax{
		newProgressiveTax(0, 150000, 0),
		newProgressiveTax(150000, 500000, 0.1),
		newProgressiveTax(500000, 1000000, 0.15),
		newProgressiveTax(1000000, 2000000, 0.2),
		newProgressiveTax(2000000, math.Inf(1), 0.35),
	}
}