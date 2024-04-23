package handlers

func CalculatePersonalDeductionTax(income float64, deduction float64) float64 {
	return income - deduction
}