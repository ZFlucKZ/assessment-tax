package dto

type AllowanceType struct {
	AllowanceType string  `json:"allowanceType"`
	Amount        float64 `json:"amount"`
}

type TaxLevel struct {
	Level string  `json:"level"`
	Tax   float64 `json:"tax"`
}

type Tax struct {
	TotalIncome float64         `json:"totalIncome"`
	Wht         float64         `json:"wht"`
	Allowances  []AllowanceType `json:"allowances"`
}

type TaxResponse struct {
	Tax      float64    `json:"tax"`
	TaxLevel []TaxLevel `json:"taxLevel"`
}

type TaxRefundResponse struct {
	TaxRefund float64    `json:"taxRefund"`
	TaxLevel  []TaxLevel `json:"taxLevel"`
}

type TaxCSV struct {
	TotalIncome float64 `json:"totalIncome" csv:"totalIncome"`
	Wht         float64 `json:"wht" csv:"wht"`
	Donation    float64 `json:"donation" csv:"donation"`
	KReceipt    float64 `json:"k-receipt" csv:"k-receipt"`
}

type Taxes struct {
	TotalIncome float64 `json:"totalIncome"`
	Tax         float64 `json:"tax"`
}

type TaxesRefund struct {
	TotalIncome float64 `json:"totalIncome"`
	TaxRefund   float64 `json:"taxRefund"`
}

type TaxCSVResponse struct {
	Taxes []interface{} `json:"taxes"`
}