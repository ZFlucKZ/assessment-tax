package models

type Deduction struct {
	Personal  float64 `json:"personal"`
	K_receipt float64 `json:"k_receipt"`
}

func InitDeduction() *Deduction {
	return &Deduction{
		Personal:  60000.0,
		K_receipt: 50000.0,
	}
}