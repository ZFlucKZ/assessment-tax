package routes

import (
	"github.com/ZFlucKZ/assessment-tax/controllers"
	"github.com/labstack/echo/v4"
)

func taxRoutes(e *echo.Echo) {
	tax := e.Group("/tax")

	tCl := controllers.TaxController{}
	tax.POST("/calculations", tCl.CalculateTaxHandler)
}