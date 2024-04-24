package handlers

import "github.com/ZFlucKZ/assessment-tax/dto"

func CalculateTotalTax(taxDetails *dto.Tax) float64 {
	// Find Allowance type Personal From TaxDetails Allowances
	personalAllowance := dto.AllowanceType{}
	for _, allowance := range taxDetails.Allowances {
		if allowance.AllowanceType == "Personal" {
			personalAllowance = allowance
			break
		}
	}

	taxDetails.TotalIncome = calculatePersonalDeductionTax(taxDetails.TotalIncome, personalAllowance.Amount)

	taxes := initProgressiveTax()

	// Calculate Total Tax
	var totalTax float64
	for _, tax := range taxes {
		totalTax += tax.calculateTax(taxDetails.TotalIncome)
	}

	// Calculate Total Tax with WHT
	totalTax = calculateTaxWithWht(totalTax, taxDetails.Wht)

	return totalTax
}