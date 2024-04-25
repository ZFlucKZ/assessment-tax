package routes

import (
	"github.com/ZFlucKZ/assessment-tax/controllers"
	"github.com/ZFlucKZ/assessment-tax/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func adminRoutes(e *echo.Echo) {
	admin := e.Group("/admin")

	admin.Use(middleware.BasicAuth(middlewares.BasicAuthMiddleware))

	aCl := controllers.AdminController{}

	admin.POST("/deductions/personal", aCl.UpdatePersonalDeduction)
}