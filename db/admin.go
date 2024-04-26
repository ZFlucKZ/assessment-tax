package db

import "database/sql"

var DB *sql.DB

func SetDatabase(p *sql.DB) {
	DB = p
}

func UpdatePersonalDeduction(amount float64) (float64, error) {
	_, err := DB.Exec("UPDATE deduction SET amount = $1, updated_at = CURRENT_TIMESTAMP WHERE deduction_type = 'Personal'", amount)
	if err != nil {
		return 0, err
	}

	return amount, nil
}

func UpdateKReceiptDeduction(amount float64) (float64, error) {
	_, err := DB.Exec("UPDATE deduction SET amount = $1, updated_at = CURRENT_TIMESTAMP WHERE deduction_type = 'KReceipt'", amount)
	if err != nil {
		return 0, err
	}

	return amount, nil
}

