package controllers

import (
	"math"
	"net/http"

	"github.com/ZFlucKZ/assessment-tax/dto"
	"github.com/ZFlucKZ/assessment-tax/handlers"
	"github.com/ZFlucKZ/assessment-tax/models"
	"github.com/labstack/echo/v4"
)

type TaxController struct{
}

func (t TaxController) CalculateTaxHandler(c echo.Context) error {
	taxDetails := new(dto.Tax)

	err := c.Bind(taxDetails);
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Err{Message: "Invalid request payload"})
	}

	if taxDetails.TotalIncome < 0 {
		return c.JSON(http.StatusBadRequest, dto.Err{Message: "Total income must be greater or equal 0"})
	}

	if taxDetails.Wht < 0 || taxDetails.Wht > taxDetails.TotalIncome {
		return c.JSON(http.StatusBadRequest, dto.Err{Message: "WHT must be greater or equal 0 and less than total income"})
	}

	deduction := models.InitDeduction()

	taxDetails.Allowances = append(taxDetails.Allowances, dto.AllowanceType{AllowanceType: "Personal", Amount: deduction.Personal})

	tax := handlers.CalculateTotalTax(taxDetails)

	if tax >= 0 {
		return c.JSON(http.StatusOK, dto.TaxResponse{Tax: tax})
	}else {
		return c.JSON(http.StatusOK, dto.TaxRefundResponse{TaxRefund: math.Abs(tax)})
	}
}