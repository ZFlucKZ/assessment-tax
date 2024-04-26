package handlers

import (
	"errors"

	"github.com/ZFlucKZ/assessment-tax/dto"
)

func CalculateTotalTax(taxDetails *dto.Tax) (float64, []dto.TaxLevel, error) {
	// Find Allowance type Personal From TaxDetails Allowances
	personalAllowance := dto.AllowanceType{}
	donationAllowance := dto.AllowanceType{}
	kReceiptAllowance := dto.AllowanceType{}

	for _, allowance := range taxDetails.Allowances {
		if allowance.Amount < 0 {
			return 0.0, nil, errors.New("Allowance amount must be greater or equal 0")
		}

		if allowance.AllowanceType == "Personal" {
			personalAllowance = allowance
		} else if allowance.AllowanceType == "donation" {
			donationAllowance = allowance
		} else if allowance.AllowanceType == "k-receipt" {
			kReceiptAllowance = allowance
		}
	}

	var err error

	taxDetails.TotalIncome, err = calculatePersonalDeductionTax(taxDetails.TotalIncome, personalAllowance.Amount)
	if err != nil {
		return 0.0, nil, errors.New("Failed to calculate personal deduction tax")
	}
	
	taxDetails.TotalIncome, err = calculateDonationDeductionTax(taxDetails.TotalIncome, donationAllowance.Amount)
	if err != nil {
		return 0.0, nil, errors.New("Failed to calculate donation deduction tax")
	}

	taxDetails.TotalIncome, err = calculateKReceiptDeductionTax(taxDetails.TotalIncome, kReceiptAllowance.Amount)
	if err != nil {
		return 0.0, nil, errors.New("Failed to calculate k-receipt deduction tax")
	}
	
	taxes := initProgressiveTax()

	// Calculate Total Tax
	var totalTax float64
	var taxLevelList []dto.TaxLevel

	for _, tax := range taxes {
		tax, taxLevel := tax.calculateTax(taxDetails.TotalIncome)
		totalTax += tax
		taxLevelList = append(taxLevelList, taxLevel)
	}

	// Calculate Total Tax with WHT
	totalTax = calculateTaxWithWht(totalTax, taxDetails.Wht)

	return totalTax, taxLevelList, nil
}