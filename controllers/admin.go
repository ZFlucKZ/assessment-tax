package controllers

import (
	"net/http"

	"github.com/ZFlucKZ/assessment-tax/db"
	"github.com/ZFlucKZ/assessment-tax/dto"
	"github.com/labstack/echo/v4"
)

type AdminController struct{}

func (a AdminController) UpdatePersonalDeduction(c echo.Context) error {
	payload := new(dto.DeductionAmount)

	err := c.Bind(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Message{Message: "Invalid request payload"})
	}

	if payload.Amount < 10000 || payload.Amount > 100000{
		return c.JSON(http.StatusBadRequest, dto.Message{Message: "Personal deduction must be between 10,000 and 100,000"})
	}

	amount, err := db.UpdatePersonalDeduction(payload.Amount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Message{Message: "Failed to update personal deduction. Please try again later."})
	}

	return c.JSON(http.StatusOK, dto.ResponsePersonal{PersonalDeduction: amount})
}

func (a AdminController) UpdateKReceiptDeduction(c echo.Context) error {
	payload := new(dto.DeductionAmount)

	err := c.Bind(payload)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.Message{Message: "Invalid request payload"})
	}

	if payload.Amount < 0 || payload.Amount > 100000{
		return c.JSON(http.StatusBadRequest, dto.Message{Message: "K-Receipt deduction must be between 0 and 100,000"})
	}

	amount, err := db.UpdateKReceiptDeduction(payload.Amount)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.Message{Message: "Failed to update K-Receipt deduction. Please try again later."})
	}

	return c.JSON(http.StatusOK, dto.ResponseKReceipt{KReceipt: amount})
}