package controllers

import (
	"net/http"

	"github.com/ZFlucKZ/assessment-tax/dto"
	"github.com/ZFlucKZ/assessment-tax/handlers"
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

	// fetch personalDeduction from database if doesn't exist use default value (60,000)
	taxDetails.Allowances = append(taxDetails.Allowances, dto.AllowanceType{AllowanceType: "Personal", Amount: 60000})

	tax := handlers.CalculateProgressiveTax(taxDetails)

	return c.JSON(http.StatusOK, dto.TaxResponse{Tax: tax})
}