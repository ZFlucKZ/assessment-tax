package handlers

func calculateTaxWithWht(tax float64, wht float64) float64 {
	return tax - wht
}