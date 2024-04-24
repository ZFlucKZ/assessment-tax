package handlers

func calculatePersonalDeductionTax(income float64, personal float64) float64 {
	return income - personal
}

func calculateDonationDeductionTax(income float64, donation float64) float64 {
	// donation max 100000
	if donation > 100000 {
		donation = 100000
	}

	return income - donation
}