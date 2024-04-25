package db

import (
	"github.com/ZFlucKZ/assessment-tax/models"
)

func  GetDeductionAmountByDeductionType(deductionType string) (models.DeductionModel, error) {
	var deduction models.DeductionModel
	err := DB.QueryRow("SELECT deduction_type, amount FROM deduction WHERE deduction_type = $1", deductionType).Scan(&deduction.Deduction, &deduction.Amount)
	if err != nil {
		return models.DeductionModel{}, err
	}

	return deduction, nil
}