package handlers

import (
	"fmt"
	"math"

	"github.com/ZFlucKZ/assessment-tax/dto"
	humanize "github.com/dustin/go-humanize"
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

func (p *ProgressiveTax) calculateTax(income float64) (float64, dto.TaxLevel) {
	level := fmt.Sprintf("%s-%s", humanize.Commaf((p.minIncome+1)), humanize.Commaf(p.maxIncome))

	if p.minIncome == 2000000 {
		level = fmt.Sprintf("%s ขึ้นไป", humanize.Commaf(p.minIncome+1))
	}

	if p.minIncome == 0 {
		level = fmt.Sprintf("%s-%s", humanize.Commaf((p.minIncome)), humanize.Commaf(p.maxIncome))	}

	if income < p.minIncome {
		return 0, dto.TaxLevel{
			Level: level,
			Tax:   0,
		}
	}

	if income > p.maxIncome {
		return (p.maxIncome - p.minIncome) * p.rate, dto.TaxLevel{
			Level: level,
			Tax:   (p.maxIncome - p.minIncome) * p.rate,
		}
	}

	return (income - p.minIncome) * p.rate, dto.TaxLevel{
		Level: level,
		Tax:   (income - p.minIncome) * p.rate,
	}
}

func initProgressiveTax() []*ProgressiveTax {
	return []*ProgressiveTax{
		newProgressiveTax(0, 150_000, 0),
		newProgressiveTax(150_000, 500_000, 0.1),
		newProgressiveTax(500_000, 1_000_000, 0.15),
		newProgressiveTax(1_000_000, 2_000_000, 0.2),
		newProgressiveTax(2_000_000, math.Inf(1), 0.35),
	}
}