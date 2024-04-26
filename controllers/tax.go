package controllers

import (
	"math"
	"net/http"

	"github.com/ZFlucKZ/assessment-tax/dto"
	"github.com/ZFlucKZ/assessment-tax/handlers"
	"github.com/gocarina/gocsv"
	"github.com/labstack/echo/v4"
)

type TaxController struct{}

func (t TaxController) CalculateTaxHandler(c echo.Context) error {
	taxDetails := new(dto.Tax)

	err := c.Bind(taxDetails);
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Message{Message: "Invalid request payload"})
	}

	if taxDetails.TotalIncome < 0 {
		return c.JSON(http.StatusBadRequest, dto.Message{Message: "Total income must be greater or equal 0"})
	}

	if taxDetails.Wht < 0 || taxDetails.Wht > taxDetails.TotalIncome {
		return c.JSON(http.StatusBadRequest, dto.Message{Message: "WHT must be greater or equal 0 and less than total income"})
	}

	tax, taxLevel, err := handlers.CalculateTotalTax(taxDetails)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Message{Message: "Failed to calculate tax. Please try again later."})
	}

	if tax >= 0 {
		return c.JSON(http.StatusOK, dto.TaxResponse{Tax: tax, TaxLevel: taxLevel})
	}else {
		return c.JSON(http.StatusOK, dto.TaxRefundResponse{TaxRefund: math.Abs(tax), TaxLevel: taxLevel})
	}
}

func (t TaxController) CalculateTaxByCsvFileHandler(c echo.Context) error {
	file, err := c.FormFile("taxFile")
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	taxesDetails := []*dto.TaxCSV{}

	err = gocsv.UnmarshalMultipartFile(&src, &taxesDetails); 
	if err != nil {
		panic(err)
	}

	taxes := dto.TaxCSVResponse{}

	for _, row := range taxesDetails {
		if row.TotalIncome < 0 {
			return c.JSON(http.StatusBadRequest, dto.Message{Message: "Total income must be greater or equal 0"})
		}

		if row.Wht < 0 || row.Wht > row.TotalIncome {
			return c.JSON(http.StatusBadRequest, dto.Message{Message: "WHT must be greater or equal 0 and less than total income"})
		}

		if row.Donation < 0 || row.KReceipt < 0 {
			return c.JSON(http.StatusBadRequest, dto.Message{Message: "Donation and K-Receipt must be greater or equal 0"})
		}

		taxDetails := dto.Tax{
			TotalIncome: row.TotalIncome,
			Wht: row.Wht,
		}

		allowances := []dto.AllowanceType{
			{
				AllowanceType: "donation",
				Amount: row.Donation,
			},
			{
				AllowanceType: "k-receipt",
				Amount: row.KReceipt,
			},
		}

		taxDetails.Allowances = allowances

		tax, _, err := handlers.CalculateTotalTax(&taxDetails)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, dto.Message{Message: "Failed to calculate tax. Please try again later."})
		}

		if tax >= 0 {
			taxes.Taxes = append(taxes.Taxes, dto.Taxes{TotalIncome: row.TotalIncome, Tax: tax})
		}else {
			taxes.Taxes = append(taxes.Taxes, dto.TaxesRefund{TotalIncome: row.TotalIncome, TaxRefund: math.Abs(tax)})
		}
	}

	return c.JSON(http.StatusOK, taxes)
}