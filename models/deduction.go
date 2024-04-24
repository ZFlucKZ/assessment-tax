package models

type Deduction struct {
	Personal float64 `json:"personal"`
}

func InitDeduction() *Deduction {
	return &Deduction{
		Personal: 60000.0,
	}
}