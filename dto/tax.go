package dto

type AllowanceType struct {
	AllowanceType string  `json:"allowanceType"`
	Amount        float64 `json:"amount"`
}

type Tax struct {
	TotalIncome float64         `json:"totalIncome"`
	Wht         float64         `json:"wht"`
	Allowances  []AllowanceType `json:"allowances"`
}

type TaxResponse struct {
	Tax float64 `json:"tax"`
}

type TaxRefundResponse struct {
	TaxRefund float64 `json:"taxRefund"`
}