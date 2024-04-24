package handlers

import (
	"errors"

	"github.com/ZFlucKZ/assessment-tax/dto"
)

func CalculateTotalTax(taxDetails *dto.Tax) (float64, error) {
	// Find Allowance type Personal From TaxDetails Allowances
	personalAllowance := dto.AllowanceType{}
	donationAllowance := dto.AllowanceType{}
	// kReceiptAllowance := dto.AllowanceType{}
	for _, allowance := range taxDetails.Allowances {
		if allowance.Amount < 0 {
			return 0, errors.New("Invalid allowance amount")
		}

		if allowance.AllowanceType == "Personal" {
			personalAllowance = allowance
		} else if allowance.AllowanceType == "donation" {
			donationAllowance = allowance
		} else if allowance.AllowanceType == "k-receipt" {
			// kReceiptAllowance := allowance
		}
	}

	taxDetails.TotalIncome = calculatePersonalDeductionTax(taxDetails.TotalIncome, personalAllowance.Amount)
	
	taxDetails.TotalIncome = calculateDonationDeductionTax(taxDetails.TotalIncome, donationAllowance.Amount)
	
	taxes := initProgressiveTax()

	// Calculate Total Tax
	var totalTax float64
	for _, tax := range taxes {
		totalTax += tax.calculateTax(taxDetails.TotalIncome)
	}

	// Calculate Total Tax with WHT
	totalTax = calculateTaxWithWht(totalTax, taxDetails.Wht)

	return totalTax, nil
}