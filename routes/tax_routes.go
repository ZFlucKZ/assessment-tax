package routes

import (
	"github.com/ZFlucKZ/assessment-tax/controllers"
	"github.com/labstack/echo/v4"
)

func taxRoutes(e *echo.Echo) {
	tax := e.Group("/tax")

	taxController := controllers.TaxController{}
	tax.POST("/calculations", taxController.CalculateTax)
}