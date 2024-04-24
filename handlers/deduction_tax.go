package handlers

func calculatePersonalDeductionTax(income float64, deduction float64) float64 {
	return income - deduction
}