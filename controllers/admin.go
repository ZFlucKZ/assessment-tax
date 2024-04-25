package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type AdminController struct{}

func (a AdminController) UpdatePersonalDeduction(c echo.Context) error {
	return c.JSON(http.StatusOK, "Update personal deduction")
}