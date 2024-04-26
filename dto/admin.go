package dto

type DeductionAmount struct {
	Amount float64 `json:"amount"`
}

type ResponsePersonal struct {
	PersonalDeduction float64 `json:"personalDeduction"`
}

type ResponseKReceipt struct {
	KReceipt float64 `json:"kReceipt"`
}